// @title Golang API
// @version 0.0.1
// @description Simple REST API using golang

// @contact.name Developers
// @contact.email dev@localhost

// @host localhost:4000
// @BasePath /
package main

import (
	"log"
	"net/http"

	"github.com/gpng/go-docker-api-boilerplate/config"
	"github.com/gpng/go-docker-api-boilerplate/connections/database"
	"github.com/gpng/go-docker-api-boilerplate/services/somesvc"
	vr "github.com/gpng/go-docker-api-boilerplate/utils/validator"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	appConfig := config.New()

	// initialise utils
	validator := vr.New()

	// initialise dependencies for service
	// postgres
	db, err := database.New(appConfig)
	if err != nil {
		log.Fatalf("Failed to initialise DB connection")
		return
	}

	someService := somesvc.New(db, validator)

	// initialise main router with basic middlewares, cors settings etc
	router := mainRouter(appConfig.Docs)

	// mount services
	router.Mount("/some", someService.Routes())

	err = http.ListenAndServe(":4000", router)
	if err != nil {
		log.Print(err)
	}
}

func mainRouter(docs bool) chi.Router {
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})
	router.Use(c.Handler)

	if docs {
		router.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/docs/index.html", http.StatusMovedPermanently)
		})
		router.Get("/docs/*", httpSwagger.Handler())
		log.Println("API docs available at /docs")
	}

	// stop crawlers
	router.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User-agent: *\nDisallow: /"))
	})

	return router
}
