package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

var logger *slog.Logger

func main() {
	// Initialize the logger
	logger = slog.New(slog.NewTextHandler(os.Stdout))

	http.HandleFunc("/healthz", healthCheck)
	http.HandleFunc("/your-endpoint", loggingMiddleware(yourHandler)) // Replace with your actual handler

	logger.Info("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logger.Fatal("Failed to start server", "error", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next(w, r)

		duration := time.Since(start)
		logger.Info("Request handled", 
			"method", r.Method,
			"path", r.URL.Path,
			"status", http.StatusText(http.StatusOK),
			"duration", duration)
	}
}

func yourHandler(w http.ResponseWriter, r *http.Request) {
	// Your handler logic here
	w.Write([]byte("Hello, World!"))
}package main

import \"fmt\"

func add(a int, b int) int {
	return a + b
}

func multiply(x int, y int) int {
	result := x * y
	return result
}

func divide(numerator float64, denominator float64) float64 {
	return numerator / denominator
}

func processSlice(items []int) []int {
	output := []int{}
	for _, i := range items {
		if i > 0 {
			output = append(output, i*2)
		}
	}
	return output
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(multiply(3, 4))
}
