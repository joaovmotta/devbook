package middlewares

import (
	"devbook-api/src/authorization"
	"devbook-api/src/responses"
	"log"
	"net/http"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if err := authorization.ValidateToken(r); err != nil {

			responses.Error(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}

func Logger(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
