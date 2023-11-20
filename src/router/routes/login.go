package routes

import (
	"net/http"
	"webapp/src/router/controllers"
)

var loginRoutes = []Route{
	{
		URI: "/",
		Method: http.MethodGet,
		Func:   controllers.LoadLoginPage,
		Auth:   false,
	},
	{
		URI: "/login",
		Method: http.MethodGet,
		Func:   controllers.LoadLoginPage,
		Auth:   false,
	},
	{
		URI: "/login",
		Method: http.MethodPost,
		Func:   controllers.Login,
		Auth:   false,
	},
}