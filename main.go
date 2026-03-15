package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"log/slog"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})
	
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello, World!")
		
		duration := time.Since(start)
		logger.Info("request", "method", r.Method, "path", r.URL.Path, "status", http.StatusOK, "duration", duration)
	})
	
	http.ListenAndServe(":8080", mux)
}package main

import (
    "log/slog"
    "net/http"
    "time"
)

func main() {
    logger := slog.New(slog.NewTextHandler(os.Stdout))

    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
    })

    http.HandleFunc("/your-endpoint", func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        // Your handler logic here

        statusCode := http.StatusOK // Change this based on your logic

        logger.Info("Request handled", "method", r.Method, "path", r.URL.Path, "status_code", statusCode, "duration", time.Since(start))
        w.WriteHeader(statusCode)
    })

    logger.Info("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        logger.Error("Server failed", "error", err)
    }
package main

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
