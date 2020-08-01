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
	userID := utils.GetVar(r, "userId")
	deleteCount, err := deleteUser(userID)

	if err != nil {
		utils.HandleInternalError(w, r, err)
		return
	}

	if deleteCount > 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func handleGetUser(w http.ResponseWriter, r *http.Request, claims core.JwtClaims) {
	userID := utils.GetVar(r, "userId")
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

func handleLogin(w http.ResponseWriter, r *http.Request) {

	// --- JSON-based authentication
	if r.Body == nil {
		w.WriteHeader(400)
		return
	}

	var user userWithPassword
	json.NewDecoder(r.Body).Decode(&user)

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
		json.NewEncoder(w).Encode(struct{ Token string }{Token: token})

		return
	}

}

func handleUpdateUser(w http.ResponseWriter, r *http.Request, claims core.JwtClaims) {
	var updatingUser core.User
	json.NewDecoder(r.Body).Decode(&updatingUser)

	userID := utils.GetVar(r, "userId")
	result, err := updateUser(userID, updatingUser)
	if err != nil {
		utils.HandleInternalError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(result)
}
