package routes

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
	Auth   bool
}

func Config(router *mux.Router) *mux.Router{
	routes := loginRoutes
	routes = append(routes, userRoutes...)
	routes = append(routes, homePageRoute)
	routes = append(routes, postRoutes...)
	routes = append(routes, LogoutRoute)

	for _, route := range routes {
		if route.Auth {
			router.HandleFunc(route.URI, middlewares.Logger(middlewares.AuthHandler(route.Func))).Methods(route.Method)
		}

		router.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/",fileServer))

	return router
}