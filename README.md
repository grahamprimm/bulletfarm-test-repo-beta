
## ## Health Check Endpoint
## Health Check Endpoint
A new health check endpoint has been added to the HTTP server:

### /healthz
- **Method:** GET
- **Response:** 200 OK
- Logs the health check request with method and path.

## Structured Logging
Structured logging has been implemented using the `log/slog` package. Each request to the main endpoint (/) is logged with the following details:
- **Method:** HTTP method used in the request
- **Path:** URL path requested
- **Status Code:** HTTP status code returned
- **Duration:** Time taken to process the request


## ### Health Check Endpoint
The application includes a health check endpoint at `/healthz` that responds with a 200 OK status and logs the request method and path.

### Logging
The application uses structured logging with the `log/slog` package to log each HTTP request. The logs include:
- Request method
- Request path
- Response status code
- Duration of the request

Example log entry:
```
INFO: Request handled method=GET path=/ status_code=200 duration=5ms
```

## API Endpoints
## API Endpoints

### Metrics
## Metrics

The application exposes a `/metrics` endpoint that provides Prometheus-style metrics. The following metrics are available:

- **request_count**: Total number of requests received by the server.
- **request_duration**: Duration of requests in seconds.
- **active_connections**: Current number of active connections to the server.

### Example Query

To retrieve metrics, make a request to the `/metrics` endpoint:

```
curl http://localhost:8080/metrics
```

The output is formatted in Prometheus text format, suitable for scraping by Prometheus.
