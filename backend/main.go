package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"systolic/controllers"
	"systolic/router"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("consumer_origin")},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Access-Control-Allow-Credentials", "Access-Control-Allow-Origin", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	mw := io.MultiWriter(os.Stdout, controllers.LogFile)
	log.SetOutput(mw)
	customLogger := log.New(controllers.LogFile, "", log.LstdFlags)
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: customLogger, NoColor: true})
	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(middleware.Timeout(60 * time.Second))

	router.SetupRoutes(r)
	http.ListenAndServe(":"+os.Getenv("webapp_port"), r)
}
