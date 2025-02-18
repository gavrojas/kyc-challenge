package validations

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func AddValidationRoutes(r chi.Router, db *gorm.DB) {
	r.Route("/validations", func(r chi.Router) {
		r.Post("/create", func(w http.ResponseWriter, r *http.Request) {
			CreateValidation(w, r, db)
		})
		r.Post("/put-file", func(w http.ResponseWriter, r *http.Request) {
			UploadImage(w, r, db)
		})
		r.Get("/result", func(w http.ResponseWriter, r *http.Request) {
			GetValidation(w, r, db)
		})

		r.Get("/get-config", func(w http.ResponseWriter, r *http.Request) {
			GetConfigurations(w, r, db)
		})
	})
}
