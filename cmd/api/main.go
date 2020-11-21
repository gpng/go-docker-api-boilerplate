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

	"github.com/gpng/go-docker-api-boilerplate/cmd/api/config"
	"github.com/gpng/go-docker-api-boilerplate/cmd/api/handlers"
	"github.com/gpng/go-docker-api-boilerplate/services/logger"
	"github.com/gpng/go-docker-api-boilerplate/services/postgres"
	"github.com/gpng/go-docker-api-boilerplate/services/validator"
	"github.com/gpng/go-docker-api-boilerplate/sqlc/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to load env vars: %v", err)
	}

	// initialise services
	l := logger.New()
	defer l.Sync()

	vr := validator.New()

	db, err := postgres.New(cfg.DbHost, cfg.DbUser, cfg.DbName, cfg.DbPassword)
	if err != nil {
		log.Fatalf("failed to initialise DB connection: %v", err)
	}

	repo := models.New(db)

	handlers := handlers.New(l, vr, db, repo)

	// initialise main router with basic middlewares, cors settings etc
	router := mainRouter(cfg.Docs, cfg.CORS)

	// mount services
	router.Mount("/", handlers.Routes())

	err = http.ListenAndServe(":4000", router)
	if err != nil {
		log.Print(err)
	}
}

func mainRouter(docs bool, useCors bool) chi.Router {
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	if useCors {
		c := cors.New(cors.Options{
			AllowedHeaders: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		})
		router.Use(c.Handler)
	}

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
