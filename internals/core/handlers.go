package core

import (
	"fmt"
	"net/http"
)

// HandleHealthcheck is a basic "Hey-I-m-alive" endpoint
func HandleHealthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Looks all good...\n")
}
