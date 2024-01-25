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
	URL       string `json: "url" validate: "required, url"`
	SHORT_URL string `json: "short_url, omitempty"`
}

type URLSaver interface {
	SaveUrl(urlToSave, short_url string) error
}

// TODO: move to config
const shortUrlLen = 6

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.With(
			slog.String("op", "handlers.url.save.New"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
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

		//short_url must be unique, if it is not, then generate a new one
		//TODO: refactor
		short_url := req.SHORT_URL
		if short_url == "" {
			for {
				short_url = util.NewRandomShortUrl(shortUrlLen)
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
		} else {
			err = urlSaver.SaveUrl(req.URL, short_url)
			if err == nil {
				render.JSON(w, r, Response{
					SHORT_URL: short_url,
					STATUS:    "OK",
				},
				)
			} else {
				render.JSON(w, r, Error("short_url already exists"))
				return
			}
		}
	}
}
