# BulletFarm Test Repo Beta

## Getting Started

To get started with the BulletFarm Test Repo, follow these instructions:

### Prerequisites
- Go 1.15 or later

### Build and Run Instructions
1. Clone the repository:
   ```bash
   git clone https://github.com/grahamprimm/bulletfarm-test-repo-beta.git
   cd bulletfarm-test-repo-beta
   ```
2. Build the application:
   ```bash
   go build -o bulletfarm
   ```
3. Run the application:
   ```bash
   ./bulletfarm
   ```

### API Documentation

#### Health Check Endpoint

- **Endpoint:** `/healthz`
- **Method:** `GET`
- **Response:** `200 OK`

**Example Request:**
```bash
curl -X GET http://localhost:8080/healthz
```

**Example Response:**
```json
{
  "status": "ok"
}
```

### Structured Logging
Structured logging has been implemented using the `log/slog` package. Each request to the main endpoint (/) is logged with the following details:
- **Method:** HTTP method used in the request
- **Path:** URL path requested
- **Status Code:** HTTP status code returned
- **Duration:** Time taken to process the request

