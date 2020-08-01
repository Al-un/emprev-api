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
	serverPort := 8000

	if envServerPort := os.Getenv("PORT"); envServerPort != "" {
		serverPortAsInt, err := strconv.Atoi(envServerPort)
		if err != nil {
			utils.ApiLogger.Fatalf("Invalid server port %v\n", envServerPort)
			return
		}

		serverPort = serverPortAsInt
	}

	router := mux.NewRouter()

	router.Use(core.AddCorsHeaders)
	router.Use(core.AddJSONHeaders)

	users.LoadEndpoints(router)
	reviews.LoadEndpoints(router)

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		m, err := route.GetMethods()
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		utils.ApiLogger.Infof("%-6s: %s\n", m, t)
		return nil
	})

	utils.ApiLogger.Infof("[Server] Starting server on port %d...", serverPort)
	utils.ApiLogger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serverPort), router))
}
