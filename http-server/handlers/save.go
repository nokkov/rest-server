package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	URL       string `json: "url" validate: "required, url"`
	SHORT_URL string `json: "short_url, omitempty"`
}

type URLSaver interface {
	SaveUrl(urlToSave, short_url string) error
}

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

		if err := validator.New().Struct(req); err != nil {
			log.With("struct validation fault: ", err)
			render.JSON(w, r, Error("failed to validate request"))
			return
		}

		//TODO: save url
	}
}
