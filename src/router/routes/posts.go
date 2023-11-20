package routes

import (
	"net/http"
	"webapp/src/router/controllers"
)

var postRoutes = []Route{
	{
		URI:    "/posts",
		Method: http.MethodPost,
		Func: controllers.CreatePost,
		Auth: true,
	},
}