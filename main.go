package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Al-un/emprev-api/internals/core"
	"github.com/Al-un/emprev-api/internals/reviews"
	"github.com/Al-un/emprev-api/internals/users"
	"github.com/Al-un/emprev-api/internals/utils"
	"github.com/gorilla/mux"
)

func main() {
	// --- Port definition
	serverPort := 8000

	if envServerPort := os.Getenv("PORT"); envServerPort != "" {
		serverPortAsInt, err := strconv.Atoi(envServerPort)
		if err != nil {
			utils.APILogger.Fatalf("Invalid server port %v\n", envServerPort)
			return
		}

		serverPort = serverPortAsInt
	}

	// --- Router setup
	router := mux.NewRouter()

	router.Use(core.AddCorsHeaders)
	router.Use(core.AddJSONHeaders)

	users.LoadEndpoints(router)
	reviews.LoadEndpoints(router)
	router.HandleFunc("/", core.HandleHealthcheck).Methods(http.MethodGet, http.MethodOptions)
	router.Use(mux.CORSMethodMiddleware(router))

	// For debugging: display all endpoints paths and methods
	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		m, err := route.GetMethods()
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		utils.APILogger.Infof("%-6s: %s\n", m, t)
		return nil
	})

	// --- One-time setup
	users.CreateRootIfNotExist()

	// --- Go!
	utils.APILogger.Infof("[Server] Starting server on port %d...", serverPort)
	utils.APILogger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), router))
}
