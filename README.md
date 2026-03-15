# bulletfarm-test-repo-beta
Test repository for bulletfarm agent operator (beta)

## Getting Started
# Getting Started

## Prerequisites
- Go 1.18+ installed on your machine.
- Git installed on your machine.

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/grahamprimm/bulletfarm-test-repo-beta.git
   cd bulletfarm-test-repo-beta
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Build and Run
1. Build the application:
   ```bash
   go build -o bulletfarm
   ```
2. Run the application:
   ```bash
   ./bulletfarm
   ```
3. The server will start on `localhost:8080`.

  { "id": 3, "name": "New Item" }
  ```

### PUT /api/items/{id}
- **Description:** Update an item by ID.
- **Request:**

  ```http
  PUT /api/items/1 HTTP/1.1
  Host: localhost:8080
  Content-Type: application/json

  {
      "name": "Updated Item"
  }
  ```
- **Response:**
  ```json
  { "id": 1, "name": "Updated Item" }
  ```

### DELETE /api/items/{id}
- **Description:** Delete an item by ID.
- **Request:**

  ```http
  DELETE /api/items/1 HTTP/1.1
  Host: localhost:8080
  ```
- **Response:**
  ```json
  { "message": "Item deleted" }
  ```

