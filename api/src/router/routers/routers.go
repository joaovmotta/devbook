package routers

import (
	"devbook-api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI               string
	Method            string
	Function          func(http.ResponseWriter, *http.Request)
	NeedAuthorization bool
}

func Configure(r *mux.Router) *mux.Router {

	routers := userRoutes
	routers = append(routers, loginRoute)

	for _, route := range routers {

		if route.NeedAuthorization {

			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authorization(route.Function))).Methods(route.Method)
		} else {

			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
