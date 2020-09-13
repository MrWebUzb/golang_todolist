package handlers

import (
	"net/http"

	"github.com/MrWebUzb/apiserver/app/utils"

	"github.com/MrWebUzb/apiserver/app/responses"
)

// NotFoundHandler route
func NotFoundHandler() http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusNotFound)
			utils.JSONSerializer(
				rw,
				responses.NewErrorResponse(http.StatusNotFound),
			)
		},
	)
}

// MethodNotAllowedHandler route
func MethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			utils.JSONSerializer(
				rw,
				responses.NewErrorResponse(http.StatusMethodNotAllowed),
			)
		},
	)
}
