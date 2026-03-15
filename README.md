
## API Endpoints
### Sample Request and Response

#### Request:
```http
GET /healthz HTTP/1.1
Host: localhost:8080
```

#### Response:
```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "status": "ok"
}
```
tory: `cd bulletfarm-test-repo-beta`
3. Build the server: `go build -o server main.go`
4. Run the server: `./server`

## API Endpoints

### Health Check Endpoint
A new health check endpoint has been added to the HTTP server:

#### /healthz
- **Method:** GET
- **Response:** 200 OK
- Logs the health check request with method and path.


A new health check endpoint has been added to the HTTP server:

### /healthz
- **Method:** GET
- **Response:** 200 OK
- Logs the health check request with method and path.

