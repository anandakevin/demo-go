# Book Management REST API

A simple REST API built with Echo framework for managing book data in memory.

## Features

- Create, Read, Update, Delete (CRUD) operations for books
- In-memory storage with unique ISBN validation
- Pagination support
- Asynchronous logging
- Built with Echo framework for high performance and minimal memory allocation
- Built-in middleware for logging and panic recovery

## How to Run

1. Make sure you have Go installed on your system
2. Clone this repository
3. Navigate to the project directory
4. Install dependencies:

```bash
make vendor
```

5. Run the application:

```bash
make server/echo
```

6. The server will start on port 8080

## 📖 API Documentation

Base URL
[http://localhost:8080](http://localhost:8080)

### 📚 Data Model

Each book contains the following fields:

```json
{
  "title": "string (required)",
  "author": "string (required)", 
  "isbn": "string (required, unique)",
  "release_date": "string (required, format: YYYY-MM-DDThh:mm:ssZ)"
}
```

🔗 API Endpoints
1. Create New Book

Method: POST
Endpoint: /books
Description: Creates a new book with unique ISBN validation
Request Body:

```json
{
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "isbn": "9780446310789",
    "release_date": "1960-07-11T00:00:00Z"
}
```

Success Response: 201 Created

```json
{
    "data": {
        "title": "To Kill a Mockingbird",
        "author": "Harper Lee",
        "isbn": "9780446310789",
        "release_date": "1960-07-11T00:00:00Z"
    },
    "status": "success"
}
```

Error Responses:

400 Bad Request: Invalid input data
409 Conflict: ISBN already exists

Example cURL:
```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "isbn": "9780446310789",
    "release_date": "1960-07-11T00:00:00Z"
}'
```

2. Get All Books (with Pagination and Sorting)

Method: GET
Endpoint: /books
Query Parameters:
- page (optional, default: 1)
- limit (optional, default: 10, max: 100)
- sort_by (optional, one of asc & desc)
- sort_order (optional, one of title, author, isbn, release_date)


Success Response: 200 OK

```json
{
    "data": [
        {
            "Title": "The Catcher in the Rye",
            "Author": "J.D. Salinger",
            "ISBN": "9780316769488",
            "ReleaseDate": "1951-07-16T00:00:00Z"
        },
        {
            "Title": "To Kill a Mockingbird",
            "Author": "Harper Lee",
            "ISBN": "9780446310789",
            "ReleaseDate": "1960-07-11T00:00:00Z"
        }
    ],
    "status": "success"
}
```

Example cURL:
```bash
# Get first page with default limit (10)
curl http://localhost:8080/books

# Get page 2 with limit 5
curl "http://localhost:8080/books?page=2&limit=5"
```

3. Get Book by ISBN

Method: GET
Endpoint: /books/{isbn}
Success Response: 200 OK

```json
{
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "isbn": "9780446310789",
    "release_date": "1960-07-11T00:00:00Z"
}
```

Error Response: 404 Not Found

Example cURL:

```bash
curl http://localhost:8080/books/978-0134190440
```

4. Update Book by ISBN

Method: PUT
Endpoint: /books/{isbn}
Request Body: Complete book object

```json
{
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "isbn": "9780446310789",
    "release_date": "1960-07-11T00:00:00Z"
}
```

Success Response: 200 OK
Error Responses:

400 Bad Request: Invalid input
404 Not Found: Book not found
409 Conflict: New ISBN conflicts with existing book

Example cURL:

```bash
curl -X PUT http://localhost:8080/books/9780134190440 \
  -H "Content-Type: application/json" \
  -d '{
        "title": "To Kill a Mockingbird",
        "author": "Harper Lee",
        "isbn": "9780446310789",
        "release_date": "1960-07-11T00:00:00Z"
    }'
```

5. Delete Book by ISBN

Method: DELETE
Endpoint: /books/{isbn}
Success Response: 204 No Content
Error Response: 404 Not Found

Example cURL:
```bash
curl -X DELETE http://localhost:8080/books/9780134190440
```

## Implementation Details

- Uses Echo framework for routing and middleware
- Thread-safe operations using sync.RWMutex
- Asynchronous logging using channels and goroutines
- In-memory storage using a map with ISBN as key
- Built-in request logging and panic recovery middleware
