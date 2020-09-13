package handlers

import (
	"net/http"

	"github.com/MrWebUzb/apiserver/app/responses"
	"github.com/MrWebUzb/apiserver/app/utils"
)

// HomeHandler ...
type HomeHandler struct{}

// NewHomeHandler ...
func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

// Index ...
func (h *HomeHandler) Index() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		response := struct {
			Ok      bool   `json:"ok"`
			Message string `json:"message"`
		}{
			Ok:      true,
			Message: "Welcome to TodoList api!",
		}

		err := utils.JSONSerializer(rw, response)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			utils.JSONSerializer(
				rw,
				responses.NewErrorResponse(http.StatusInternalServerError),
			)
		}
	})
}
