package main

import (
    "log"
    "net/http"
)

func main() {
    config, err := LoadConfig("./config.yaml")
    if err != nil {
        log.Fatalf("error loading config: %v", err)
    }
    
    // Example of using the config
    log.Printf("Starting server on port %s with log level %s", config.ServerPort, config.LogLevel)
    
    http.ListenAndServe(":" + config.ServerPort, nil) // Simplified for example
}