package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) != 0 {
		return value
	}
	return defaultValue
}

func main() {
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
	port := ":" + getEnv("PORT", "8080")
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
