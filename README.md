
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

## Configuration
## Configuration

The application can be configured using environment variables and a config.yaml file. Below are the available configuration options:

- **ServerPort**: (string) Port for the server to listen on. Can be set via the `SERVER_PORT` environment variable.
- **LogLevel**: (string) Log level (e.g., DEBUG, INFO, WARN, ERROR). Can be set via the `LOG_LEVEL` environment variable.
- **Timeout**: (int) Timeout value in seconds. Can be set via the `TIMEOUT` environment variable.
- **FeatureFlag**: (bool) A feature toggle. Can be set via the `FEATURE_FLAG` environment variable.

## ## Configuration
The application supports configuration via environment variables and a YAML file. Below are the configuration options:

### Configuration Options
- `SERVER_PORT`: The port that the server listens on. Default is `8080`.
- `LOG_LEVEL`: The logging level. Default is `INFO`.
- `TIMEOUT`: Timeout value for requests in seconds. Default is `30`.
- `FEATURE_X`: Enable or disable feature X. Default is `false`.
- `FEATURE_Y`: Enable or disable feature Y. Default is `false`.

### Example config.yaml
```yaml
server_port: "8080"
log_level: "INFO"
timeout: 30
feature_flags:
  feature_x: true
  feature_y: false
```
