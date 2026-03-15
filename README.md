# bulletfarm-test-repo-beta
Test repository for bulletfarm agent operator (beta)

## ## Structured Logging
This project now includes structured logging in the Go HTTP server implemented in `main.go`.\
\
### Features:\
- Each HTTP request is logged with the following details: method, path, status code, and duration.\
- A health check endpoint is available at `/healthz` to verify the server's status.\
\
### Usage:\
- Run the server and access the health check endpoint to see the logging in action.
