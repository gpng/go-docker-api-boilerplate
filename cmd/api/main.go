package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"

	"github.com/gpng/go-docker-api-boilerplate/pkg/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()

	// routes
	apiRouter.HandleFunc("/", controllers.HelloWorld).Methods("GET")

	// start server
	port := os.Getenv("PORT") // get port from .env if declared
	if port == "" {
		port = "5000"
	}

	fmt.Println("Listening on port", port)

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(router)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		fmt.Print(err)
	}

}
