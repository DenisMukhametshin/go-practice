package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"hello-http/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		body := map[string]string{"status": "ok"}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBody)
	})

	err = http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		log.Fatal(err)
	}
}
