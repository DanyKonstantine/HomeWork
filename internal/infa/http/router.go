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
		apiRouter.Post(
			"/",
			eventController.AddNewEvent(),
		)
		apiRouter.Put(
			"/",
			eventController.UppdateEvent(),
		)
		apiRouter.Delete(
			"/{id}",
			eventController.DeletEvent(),
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
		apiRouter.Post(
			"/",
			personControler.AddNewPeson(),
		)
		apiRouter.Put(
			"/",
			personControler.UppdatePeson(),
		)
		apiRouter.Delete(
			"/{id}",
			personControler.DeletePerson(),
		)
	})
}
