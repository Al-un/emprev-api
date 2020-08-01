package users

import (
	"net/http"

	"github.com/Al-un/emprev-api/internals/core"
	"github.com/gorilla/mux"
)

func LoadEndpoints(router *mux.Router) {
	router.HandleFunc("/users/login/", handleLogin).Methods(http.MethodPost, http.MethodOptions)
	router.Handle("/users/", core.DoIfAdmin(handleListUsers)).Methods(http.MethodGet, http.MethodOptions)
	router.Handle("/users/", core.DoIfAdmin(handleCreateUser)).Methods(http.MethodPost, http.MethodOptions)
	router.Handle("/users/{userId}/", core.DoIfLogged(handleGetUser)).Methods(http.MethodGet, http.MethodOptions)
	router.Handle("/users/{userId}/", core.DoIfAdmin(handleUpdateUser)).Methods(http.MethodPut, http.MethodOptions)
	router.Handle("/users/{userId}/", core.DoIfAdmin(handleDeleteUser)).Methods(http.MethodDelete, http.MethodOptions)
}
