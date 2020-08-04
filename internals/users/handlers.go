package users

import (
	"encoding/json"
	"net/http"

	"github.com/Al-un/emprev-api/internals/core"
	"github.com/Al-un/emprev-api/internals/utils"
)

func handleCreateUser(w http.ResponseWriter, r *http.Request, claims core.JwtClaims) {
	var toCreateUser userWithPassword
	json.NewDecoder(r.Body).Decode(&toCreateUser)

	result, err := createUser(toCreateUser)
	if err != nil {
		utils.HandleInternalError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request, claims core.JwtClaims) {
	userID := utils.GetVar(r, "userID")
	deleteCount, err := deleteUser(userID)

	if err != nil {
		utils.HandleInternalError(w, r, err)
		return
	}

	// Not relevant with soft-deletion as as-of Aug-2020, there is no update
	// count returned from the DAO
	if deleteCount > 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request, claims core.JwtClaims) {
	userID := utils.GetVar(r, "userID")
	user, err := getUser(userID)
	if err != nil {
		utils.HandleInternalError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func handleListUsers(w http.ResponseWriter, r *http.Request, claims core.JwtClaims) {
	users, err := listUsers()
	if err != nil {
		utils.HandleInternalError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// handleLogin handles a login request. For the moment, only authentication with
// the payload in JSON format (JSON-based authentication) is supported. In the
// future, other form of authentication, e.g. BASIC, might be required
func handleLogin(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		w.WriteHeader(400)
		return
	}

	var user userWithPassword
	json.NewDecoder(r.Body).Decode(&user)

	// For the moment, the client checks that username and password are not
	// empty. The server does not handle the check when one of the field is
	// missing
	if user.Username != "" && user.Password != "" {
		loggedUser, err := findActiveUsernamePassword(user.Username, user.Password)
		if err != nil {
			utils.HandleInternalError(w, r, err)
			return
		}

		token, err := core.GenerateJWT(*loggedUser)
		if err != nil {
			utils.HandleInternalError(w, r, err)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Token string `json:"token"`
		}{Token: token})

		return
	}
}

func handleUpdateUser(w http.ResponseWriter, r *http.Request, claims core.JwtClaims) {
	var updatingUser core.User
	json.NewDecoder(r.Body).Decode(&updatingUser)

	userID := utils.GetVar(r, "userID")
	result, err := updateUser(userID, updatingUser)
	if err != nil {
		utils.HandleInternalError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(result)
}
