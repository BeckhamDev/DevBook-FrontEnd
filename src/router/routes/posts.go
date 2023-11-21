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
	{
		URI: "/posts/{postID}/like",
		Method: http.MethodPost,
		Func: controllers.LikePost,
		Auth: true,
	},
	{
		URI: "/posts/{postID}/unlike",
		Method: http.MethodPost,
		Func: controllers.UnlikePost,
		Auth: true,
	},
	{
		URI: "/posts/{postID}/edit",
		Method: http.MethodGet,
		Func: controllers.LoadPostEditPage,
		Auth: true,
	},
	{
		URI: "/posts/{postID}",
		Method: http.MethodPut,
		Func: controllers.UpdatePost,
		Auth: true,
	},
	{
		URI: "/posts/{postID}",
		Method: http.MethodDelete,
		Func: controllers.DeletePost,
		Auth: true,
	},
}