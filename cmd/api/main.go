package main

import (
	"log"
	"net/http"
	"time"

	"github.com/troymcgahey/go-travel-service/internal/handlers"
	"github.com/troymcgahey/go-travel-service/internal/service"
)

func main() {
	mux := http.NewServeMux()

	todoService := service.NewTodoService()
	bookingHandler := handlers.NewTodoHandler(todoService)

	mux.HandleFunc("GET /health", handlers.HealthCheck)
	mux.HandleFunc("GET /todos", bookingHandler.ListTools)
	mux.HandleFunc("POST /booking", bookingHandler.CreateBooking)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      loggingMiddleware(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Server listening on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}
