package main

import (
    "log/slog"
    "net/http"
    "os"
    "time"
)

func main() {
    logger := slog.New(slog.NewTextHandler(os.Stdout))

    // Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	
    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
        logger.Info("Health check request", "method", r.Method, "path", r.URL.Path)
    })

    // Main endpoint
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        statusCode := http.StatusOK
        // Here you can add your logic for handling requests
        w.WriteHeader(statusCode)
        duration := time.Since(start)
        logger.Info("Request handled", "method", r.Method, "path", r.URL.Path, "status_code", statusCode, "duration", duration)
    })

    logger.Info("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        logger.Error("Server failed", "error", err)
    }
}