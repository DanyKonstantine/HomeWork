package http

import (
	"HomeWork/internal/infa/http/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Router(eventController *controllers.EventController, personControler *controllers.PerconControler) http.Handler {
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

			apiRouter.Group(func(apiRouter chi.Router) {
				AddEventRoutes(&apiRouter, eventController)
				AddPersonRoutes(&apiRouter, personControler)

				apiRouter.Handle("/*", NotFoundJSON())
			})
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

	})
	(*router).Route("/persononevent", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/{id}",
			eventController.PersonOnEvent(),
		)
	})
	(*router).Route("/newevent", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/",
			eventController.AddNewEvent(),
		)
	})
	(*router).Route("/uppdateevent", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/{id}",
			eventController.UppdateEvent(),
		)
	})
}
func AddPersonRoutes(router *chi.Router, personControler *controllers.PerconControler) {
	(*router).Route("/person", func(apiRouter chi.Router) {
		apiRouter.Get(
			"/",
			personControler.FindAllPer(),
		)
		apiRouter.Get(
			"/{id}",
			personControler.FindOnePer(),
		)
	})
	(*router).Route("/newperson", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/",
			personControler.AddNewPeson(),
		)
	})
	(*router).Route("/uppdateperson", func(apiRouter chi.Router) {
		apiRouter.Post(
			"/{id}",
			personControler.UppdatePeson(),
		)
	})
}
