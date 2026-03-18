package main

import (
    "log/slog"
    "net/http"
    "os"
    "sync"
    "time"
    "expvar"
)

var (
    requestCount = expvar.NewInt("request_count")
    requestDuration = expvar.NewFloat("request_duration")
    activeConnections = expvar.NewInt("active_connections")
)

func main() {
    logger := slog.New(slog.NewTextHandler(os.Stdout))

    // Health check endpoint
    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("OK"))
        logger.Info("Health check request", "method", r.Method, "path", r.URL.Path)
    })

    // Metrics endpoint
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        // Set the Content-Type header for Prometheus
        w.Header().Set("Content-Type", "text/plain")
        // Write the metrics as Prometheus text format
        expvar.Do(func(mat expvar.KeyValue) {
            // gather the metrics
            if _, err := w.Write([]byte(mat.String())); err != nil {
                logger.Error("Error writing metrics", "error", err)
            }
        })
    })

    // Main endpoint
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        statusCode := http.StatusOK
        requestCount.Add(1) // Increment request count
        activeConnections.Add(1) // Increment active connections
        defer activeConnections.Add(-1) // Decrement active connections on finish

        // Here you can add your logic for handling requests
        w.WriteHeader(statusCode)
        duration := time.Since(start).Seconds()
        requestDuration.Set(duration) // Set request duration
        logger.Info("Request handled", "method", r.Method, "path", r.URL.Path, "status_code", statusCode, "duration", duration)
    })

    logger.Info("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        logger.Error("Server failed", "error", err)
    }
}