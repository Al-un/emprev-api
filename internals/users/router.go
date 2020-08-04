package users

import (
	"net/http"

	"github.com/Al-un/emprev-api/internals/core"
	"github.com/gorilla/mux"
)

// LoadEndpoints maps the different "users" module handlers to an endpoint
// path and method.
func LoadEndpoints(router *mux.Router) {
	router.HandleFunc("/users/login/", handleLogin).Methods(http.MethodPost, http.MethodOptions)
	// Temporary allow all logged users to fetch the user list. This has to be removed when the username
	// will be attached to the review response
	router.Handle("/users/", core.DoIfLogged(handleListUsers)).Methods(http.MethodGet, http.MethodOptions)
	router.Handle("/users/", core.DoIfAdmin(handleCreateUser)).Methods(http.MethodPost, http.MethodOptions)
	// Allow GET a single user to all logged user.
	router.Handle("/users/{userID}/", core.DoIfLogged(handleGetUser)).Methods(http.MethodGet, http.MethodOptions)
	router.Handle("/users/{userID}/", core.DoIfAdmin(handleUpdateUser)).Methods(http.MethodPut, http.MethodOptions)
	router.Handle("/users/{userID}/", core.DoIfAdmin(handleDeleteUser)).Methods(http.MethodDelete, http.MethodOptions)
}
