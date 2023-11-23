package routes

import (
	"net/http"
	"webapp/src/router/controllers"
)

var userRoutes = []Route{
	{
		URI:    "/create_user",
		Method: http.MethodGet,
		Func:   controllers.LoadCreateUserPage,
		Auth: false,
	},
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func:   controllers.CreateUser,
		Auth: false,
	},
	{
		URI:    "/search_users",
		Method: http.MethodGet,
		Func:   controllers.LoadUsersPage,
		Auth: true,
	},
	{
		URI:    "/user/{userID}",
		Method: http.MethodGet,
		Func:   controllers.LoadUserProfile,
		Auth: true,
	},
}