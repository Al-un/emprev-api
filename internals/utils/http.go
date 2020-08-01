package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetVar fetch the variable defined in the route.
//
// Such method can be framework-dependent.
func GetVar(r *http.Request, varName string) string {
	return mux.Vars(r)[varName]
}

func HandleInternalError(w http.ResponseWriter, r *http.Request, err error) {
	ApiLogger.Infof("Error when handling %s: %v\n", r.URL.Path, err)

	errorText := err.Error()
	errorMsg := struct {
		Message string
	}{
		Message: errorText,
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorMsg)
}
