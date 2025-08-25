package routers

import (
	"devbook-api/src/controllers"
	"net/http"
)

var loginRoute = Route{
	URI:               "/login",
	Method:            http.MethodPost,
	Function:          controllers.Login,
	NeedAuthorization: false,
}
