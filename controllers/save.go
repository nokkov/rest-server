package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"

	"rest_server/util"
)

// TODO: use protobuf instead of json
type Request struct {
	URL string `json: "url" validate: "required, url"`
}

type URLSaver interface {
	SaveUrl(urlToSave, shortUrl string) error
}

// TODO: move to config
const shortUrlLen = 6

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.With(
			slog.String("op", "handlers.url.save.New"),
			slog.String("reqId", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, Error("failed to decode request"))
			return
		}

		log.Info("request body decoded", slog.Any("request", req))
		//TODO: error handler (switch on validator err)
		if err := validator.New().Struct(req); err != nil {
			log.With("struct validation fault: ", err)
			render.JSON(w, r, Error("failed to validate request"))
			return
		}

		for {
			short_url := util.NewRandomShortUrl(shortUrlLen)
			err = urlSaver.SaveUrl(req.URL, short_url)
			if err == nil {
				render.JSON(w, r, Response{
					SHORT_URL: short_url,
					STATUS:    "OK",
				},
				)
				return
			}
		}
	}
}
