package http

import (
	"github.com/Cry-coder/smpl_srvr/internal/infra/http/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Router(eventController *controllers.EventController) http.Handler {
	router := chi.NewRouter()

	// Health
	router.Group(func(healthRouter chi.Router) {
		healthRouter.Use(middleware.RedirectSlashes)

		healthRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())

			healthRouter.Handle("/*", NotFoundJSON())
		})
	})

	router.Group(func(apiRouter chi.Router) {
		apiRouter.Use(middleware.RedirectSlashes)

		apiRouter.Route("/v1", func(apiRouter chi.Router) {
			AddEventRoutes(&apiRouter, eventController)

			apiRouter.Handle("/*", NotFoundJSON())
		})
	})

	return router
}

func AddEventRoutes(router *chi.Router, eventController *controllers.EventController) {
	(*router).Route("/events", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			eventController.FindAll(),
		)
		apiRouter.Get(
			"/{id}",
			eventController.FindOne(),
		)
		apiRouter.Delete(
			"/{id}",
			eventController.Delete(),
		)
		apiRouter.Post(
			"/",
			eventController.Create(),
		)
		apiRouter.Put(
			"/",
			eventController.Update(),
		)
	})
}
