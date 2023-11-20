package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r:= mux.NewRouter()

	return routes.Config(r)
}