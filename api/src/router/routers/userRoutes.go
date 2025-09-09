package routers

import (
	"devbook-api/src/controllers"
	"net/http"
)

var userRoutes = []Route{

	{
		URI:               "/users",
		Method:            http.MethodPost,
		Function:          controllers.CreateUser,
		NeedAuthorization: false,
	},
	{
		URI:               "/users",
		Method:            http.MethodGet,
		Function:          controllers.FindUsers,
		NeedAuthorization: true,
	},
	{
		URI:               "/users/{userId}",
		Method:            http.MethodGet,
		Function:          controllers.FindUserById,
		NeedAuthorization: true,
	},
	{
		URI:               "/users/{userId}",
		Method:            http.MethodPut,
		Function:          controllers.UpdateUser,
		NeedAuthorization: true,
	},
	{
		URI:               "/users/{userId}",
		Method:            http.MethodDelete,
		Function:          controllers.DeleteUser,
		NeedAuthorization: true,
	},
}
