package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetVar fetch the variable defined in the route. `varName` is case
// sensitive. This method requires a Gorilla router.
func GetVar(r *http.Request, varName string) string {
	return mux.Vars(r)[varName]
}

// HandleInternalError provides a convenient way to handle internal error
// by returning a 500 error.
//
// TODO: Not all errors are internal errors. Some errors are 4xx as some
// parameters in the request can be incorrect
func HandleInternalError(w http.ResponseWriter, r *http.Request, err error) {
	APILogger.Infof("Error when handling %s: %v\n", r.URL.Path, err)

	errorText := err.Error()
	errorMsg := struct {
		Message string
	}{
		Message: errorText,
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(errorMsg)
}
