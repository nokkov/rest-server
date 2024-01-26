package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type UrlGetter interface {
	GetUrl(searchedUrl string) (string, error)
}

type GetResponse struct {
	URL string `json: "url"`
}

func Get(log *slog.Logger, urlGetter UrlGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := middleware.GetReqID(r.Context())
		log = log.With(
			slog.String("op", "handlers.url.get.Get"),
			slog.String("reqId", reqID),
		)
		//TODO:  put field names in constant
		shortURL := chi.URLParam(r, "short_url")
		if shortURL == "" {
			log.Error("failed to handle get request: short_url param lost")
			render.JSON(w, r, Error("short_url param lost"))
			return
		}

		url, err := urlGetter.GetUrl(shortURL)
		if err != nil {
			log.Error("GetURL: " + err.Error())
			w.WriteHeader(400)
			return
		}

		render.JSON(w, r, GetResponse{URL: url})
	}
}
