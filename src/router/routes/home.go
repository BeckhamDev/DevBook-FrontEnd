package routes

import (
	"net/http"
	"webapp/src/router/controllers"
)

var homePageRoute = Route{
	URI:    "/home",
	Method: http.MethodGet,
	Func:   controllers.LoadHomePage,
	Auth: true,
}