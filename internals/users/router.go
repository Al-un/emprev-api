package users

import (
	"github.com/Al-un/emprev-api/internals/core"
	"github.com/gorilla/mux"
)

func LoadEndpoints(router *mux.Router) {
	router.HandleFunc("/users/login", handleLogin).Methods("POST")
	router.Handle("/users", core.DoIfAdmin(handleListUsers)).Methods("GET")
	router.Handle("/users", core.DoIfAdmin(handleCreateUser)).Methods("POST")
	router.Handle("/users/{userId}", core.DoIfLogged(handleGetUser)).Methods("GET")
	router.Handle("/users/{userId}", core.DoIfAdmin(handleUpdateUser)).Methods("PUT")
	router.Handle("/users/{userId}", core.DoIfAdmin(handleDeleteUser)).Methods("DELETE")
}
