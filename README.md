# go-shorty

A simple URL shortener service written in Go to understand how go handles http requests.

## Features

- Shortens URLs using truncated SHA-256 hashes
- In-memory storage with LRU eviction
- Optional Redis storage for persistence
- RESTful API built with [chi](https://github.com/go-chi/chi)

## Getting Started

### Prerequisites

- Go 1.20+
- (Optional) Redis running locally on `localhost:6379`

### Running the App

```bash
go run ./cmd/main.go
```

### API Endpoints

#### `POST /shorten`

Shortens a given URL.

- **Request Body (form data):**
  - `url`: The original URL to shorten

- **Response:**
  - `200 OK`: Returns the shortened URL code
  - `500 Internal Server Error`: If storage fails

**Example:**

```bash
curl -X POST -d "url=https://example.com" http://localhost:8080/shorten
```