package main

import (
	"log/slog"
	"net/http"
	"time"
)

// healthCheckHandler returns a simple health check response.
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"status\": \"healthy\"}"))
}

// loggingMiddleware logs details of each HTTP request.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		statusCode := http.StatusOK // Placeholder for actual status code
		logger := slog.New(slog.NewTextHandler(os.Stdout)) // Initialize the slog logger
		logger.Info("HTTP request",
			"method", r.Method,
			"path", r.URL.Path,
			"status_code", statusCode,
			"duration", duration.Seconds(),
		)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthCheckHandler)
	
	http.ListenAndServe(":8080", loggingMiddleware(mux))
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
