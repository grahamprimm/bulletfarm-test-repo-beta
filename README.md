# bulletfarm-test-repo-beta
Test repository for bulletfarm agent operator (beta)

## ## Changes Made
- Added structured logging to the Go HTTP server in `main.go` using the `log/slog` package.
- Implemented logging for each request, recording the method, path, status code, and duration.
- Created a new health check endpoint at `/healthz`.
