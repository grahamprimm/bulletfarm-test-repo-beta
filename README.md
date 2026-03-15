# Bulletfarm Test Repo Beta

This repository contains a Go server for managing items. Follow the instructions below to get started with the API and understand the available endpoints.

## Getting Started

## Prerequisites
- Go (version 1.16 or later)
- Git

## Installation
1. Clone the repository:
   ```
   git clone https://github.com/grahamprimm/bulletfarm-test-repo-beta.git
   cd bulletfarm-test-repo-beta
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```

## Building the Server
To build the server, run:
```bash
go build -o bulletfarm-server main.go
```

## Running the Server
To run the server, execute:
```bash
./bulletfarm-server
```

# API Documentation

## Endpoints

### GET /api/items
#### Request Example:
```bash
curl -X GET http://localhost:8080/api/items
```

#### Response Example:
```json
[
  {
    "id": 1,
    "name": "Item 1",
    "description": "Description of Item 1"
  }
]
```

### POST /api/items
#### Request Example:
```bash
curl -X POST http://localhost:8080/api/items -H 'Content-Type: application/json' -d '{"name": "New Item", "description": "Description of New Item"}'
```

#### Response Example:
```json
{
  "id": 2,
  "name": "New Item",
  "description": "Description of New Item"
}
```

### PUT /api/items/{id}
#### Request Example:
```bash
curl -X PUT http://localhost:8080/api/items/1 -H 'Content-Type: application/json' -d '{"name": "Updated Item", "description": "Updated Description"}'
```

#### Response Example:
```json
{
  "id": 1,
  "name": "Updated Item",
  "description": "Updated Description"
}
```

### DELETE /api/items/{id}
#### Request Example:
```bash
curl -X DELETE http://localhost:8080/api/items/1
```

#### Response Example:
```json
{
  "message": "Item deleted successfully"
}
```# bulletfarm-test-repo-beta
Test repository for bulletfarm agent operator (beta)

## Getting Started
# Getting Started

## Prerequisites
- Go (version 1.16 or later)
- Git

## Installation
1. Clone the repository:
   ```
   git clone https://github.com/grahamprimm/bulletfarm-test-repo-beta.git
   cd bulletfarm-test-repo-beta
   ```
2. Install dependencies:
   ```
   go mod tidy
   ```

## Building the Server
To build the server, run:
```bash
go build -o bulletfarm-server main.go
```

## Running the Server
To run the server, execute:
```bash
./bulletfarm-server
```
Response Example:
```json
{
  "id": 2,
  "name": "New Item",
  "description": "Description of New Item"
}
```

### PUT /api/items/{id}
#### Request Example:
```bash
curl -X PUT http://localhost:8080/api/items/1 -H 'Content-Type: application/json' -d '{"name": "Updated Item", "description": "Updated Description"}'
```

#### Response Example:
```json
{
  "id": 1,
  "name": "Updated Item",
  "description": "Updated Description"
}
```

### DELETE /api/items/{id}
#### Request Example:
```bash
curl -X DELETE http://localhost:8080/api/items/1
```

#### Response Example:
```json
{
  "message": "Item deleted successfully"
}
```

## ## New Features
- Added structured logging to the Go HTTP server using the standard log/slog package.
- Each request is now logged with method, path, status code, and duration.
- Created a health check endpoint at /healthz. 


## Usage
### Usage

To run the server:

```bash
 go run main.go
```

Health check endpoint is available at `/healthz`. Each HTTP request will be logged with method, path, status code, and duration.

## ## Logging
This project implements structured logging for the Go HTTP server using the log/slog package.

### Features:
- Logs each request with method, path, status code, and duration.
- Health check endpoint available at `/healthz`, returning a 200 OK status when the server is healthy.
