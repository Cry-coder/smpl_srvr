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
			apiRouter.Put("/login", eventController.LoginPutHandler())
			apiRouter.Get("/login", eventController.LoginHandler())
			apiRouter.Handle("/*", NotFoundJSON())
		})
	})

	return router
}

func AddEventRoutes(router *chi.Router, eventController *controllers.EventController) {
	(*router).Route("/user", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/question/{id}",
			eventController.AuthRequired(eventController.OneQuestion()),
		)
		apiRouter.Get(
			"/profile",
			eventController.AuthRequired(eventController.FindOne()),
		)
		apiRouter.Post(
			"/cr",
			eventController.AuthRequired(eventController.CreateQuestion()),
		)
		apiRouter.Put(
			"/update/profile",
			eventController.AuthRequired(eventController.Update()),
		)
		apiRouter.Put(
			"/update/question",
			eventController.AuthRequired(eventController.UpdateQuestion()),
		)
		apiRouter.Get(
			"/logout",
			eventController.AuthRequired(eventController.LogOut()),
		)

	})
	(*router).Route("/admin", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/all",
			eventController.AdminAuth(eventController.FindAll()),
		)
		apiRouter.Get(
			"/one/{id}",
			eventController.AdminAuth(eventController.FindOneAdmin()),
		)
		apiRouter.Get(
			"/questions",
			eventController.AdminAuth(eventController.FindAllQuestions()),
		)
		apiRouter.Get(
			"/question/{id}",
			eventController.AdminAuth(eventController.FindOneQuestionAdmin()),
		)
		apiRouter.Get(
			"/profile",
			eventController.AdminAuth(eventController.FindOne()),
		)
		apiRouter.Post(
			"/createuser",
			eventController.AdminAuth(eventController.UserSignUp()),
		)
		apiRouter.Post(
			"/create/admin",
			eventController.AdminSignUp(),
		)
		apiRouter.Delete(
			"/delete/{id}",
			eventController.AdminAuth(eventController.Delete()),
		)
		apiRouter.Delete(
			"/delete/question/{id}",
			eventController.AdminAuth(eventController.DeleteQuestion()),
		)
		apiRouter.Put(
			"/update/question",
			eventController.AdminAuth(eventController.AdminUpdateQuestion()),
		)
		apiRouter.Get(
			"/logout",
			eventController.AdminAuth(eventController.LogOut()),
		)

	})
}
