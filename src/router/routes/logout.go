package routes

import (
	"net/http"
	"webapp/src/router/controllers"
)

var LogoutRoute = Route{
	URI:    "/logout",
	Method: http.MethodGet,
	Func: controllers.Logout,
	Auth: true,
}