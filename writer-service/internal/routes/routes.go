package routes

import (
    "github.com/go-chi/chi/v5"
    "writer/internal/controller"
)

func RegisterRoutes(
    r *chi.Mux,
    writerCtrl *controller.WriterController,
) {
    r.Route("/writers", func(r chi.Router) {
		r.Get("/", writerCtrl.GetAll)
		r.Get("/{id}", writerCtrl.GetByID)

		r.Post("/", writerCtrl.Upsert)

		r.Put("/{id}", writerCtrl.Upsert)

		r.Patch("/{id}", writerCtrl.PatchUpdate)
    })
}
