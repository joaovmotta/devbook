package routers

import (
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

		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
