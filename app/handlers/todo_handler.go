package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"

	"github.com/MrWebUzb/apiserver/app/entities"
	"github.com/MrWebUzb/apiserver/app/store"

	"github.com/MrWebUzb/apiserver/app/responses"
	"github.com/MrWebUzb/apiserver/app/utils"
)

// TodoHandler ...
type TodoHandler struct {
	log *logrus.Logger
	db  *sql.DB
}

// NewTodoHandler ...
func NewTodoHandler(log *logrus.Logger, db *sql.DB) *TodoHandler {
	return &TodoHandler{
		log: log,
		db:  db,
	}
}

// Index ...
func (h *TodoHandler) Index() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		todoStore := store.NewTodoStore(h.db)

		todos, err := todoStore.GetAll()

		if err != nil {
			h.log.Warnf("Error fetching data: %+v", err)

			rw.WriteHeader(http.StatusInternalServerError)
			utils.JSONSerializer(
				rw,
				responses.NewErrorResponse(http.StatusInternalServerError),
			)
			return
		}

		response := struct {
			Ok    bool                   `json:"ok"`
			Todos []*entities.TodoEntity `json:"todos"`
		}{
			Ok:    true,
			Todos: todos,
		}

		err = utils.JSONSerializer(rw, response)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			utils.JSONSerializer(
				rw,
				responses.NewErrorResponse(http.StatusInternalServerError),
			)
			return
		}
	})
}

// View ...
func (h *TodoHandler) View() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			utils.JSONSerializer(
				rw,
				responses.NewErrorResponse(http.StatusBadRequest),
			)
			return
		}

		todoStore := store.NewTodoStore(h.db)

		todo, err := todoStore.FindByID(id)

		if err != nil {
			h.log.Warnf("Error finding todo: %+v", err)
			if err == sql.ErrNoRows {
				rw.WriteHeader(http.StatusNotFound)
				utils.JSONSerializer(
					rw,
					responses.NewErrorResponse(http.StatusNotFound),
				)
				return
			}
			rw.WriteHeader(http.StatusBadRequest)
			utils.JSONSerializer(
				rw,
				responses.NewErrorResponse(http.StatusBadRequest),
			)
			return
		}

		response := struct {
			Ok   bool                 `json:"ok"`
			Todo *entities.TodoEntity `json:"todo"`
		}{
			Ok:   true,
			Todo: todo,
		}

		err = utils.JSONSerializer(rw, response)

		if err != nil {
			h.log.Warnf("Error serializing structure: %+v", err)
			rw.WriteHeader(http.StatusInternalServerError)
			utils.JSONSerializer(
				rw,
				responses.NewErrorResponse(http.StatusInternalServerError),
			)
			return
		}
	})
}
