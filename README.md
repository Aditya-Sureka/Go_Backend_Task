# Go User API

A small REST API for managing users built with Go, Fiber, PostgreSQL, and SQLC.

## Features

- Create, read, update, delete, and list users
- PostgreSQL-backed persistence
- SQLC-generated database access layer
- Request validation with `go-playground/validator`
- Structured logging with Zap
- Request ID and request logging middleware

## Tech Stack

- Go 1.26.3
- Fiber v2
- PostgreSQL
- SQLC
- Zap logger

## Project Structure

```text
cmd/server        Application entrypoint
config            Database connection setup
db/migrations     SQL schema files
db/query          SQLC query definitions
db/sqlc           Generated SQLC code
internal/handler   HTTP handlers
internal/middleware  Request logging and request ID middleware
internal/models    API request models
internal/repository  Data access layer
internal/routes    Route registration
internal/service   Business logic
```

## Prerequisites

- Go 1.26.3 or newer
- PostgreSQL 13+ recommended
- `sqlc` installed if you need to regenerate database code

## Configuration

Create a `.env` file in the project root:

```env
DB_URL=postgres://username:password@localhost:5432/userdb?sslmode=disable
```

`DB_URL` is the only runtime environment variable currently used by the application.

## Database Setup

The schema is defined in `db/migrations/000001_create_users.up.sql` and contains a single `users` table:

- `id` serial primary key
- `name` text, required
- `dob` date, required

Apply the migration to your PostgreSQL database before starting the API.

If you update files in `db/query` or `db/migrations`, regenerate SQLC output with:

```bash
sqlc generate
```

## Run Locally

1. Install dependencies:

```bash
go mod download
```

2. Start the server:

```bash
go run ./cmd/server
```

The API listens on port `3000`.

## Docker

Build the image:

```bash
docker build -t go-user-api -f dockerfile .
```

Run the container:

```bash
docker run --rm -p 3000:3000 --env-file .env go-user-api
```

## API Endpoints

### Health Check

`GET /`

Response:

```text
API Running
```

### Create User

`POST /users`

Request body:

```json
{
  "name": "Jane Doe",
  "dob": "1995-08-15"
}
```

Response:

```json
{
  "id": 1,
  "name": "Jane Doe",
  "dob": "1995-08-15"
}
```

### Get User

`GET /users/:id`

Response:

```json
{
  "id": 1,
  "name": "Jane Doe",
  "dob": "1995-08-15",
  "age": 30
}
```

### List Users

`GET /users`

Response:

```json
[
  {
    "id": 1,
    "name": "Jane Doe",
    "dob": "1995-08-15",
    "age": 30
  }
]
```

### Update User

`PUT /users/:id`

Request body:

```json
{
  "name": "Jane Smith",
  "dob": "1995-08-15"
}
```

Response:

```json
{
  "id": 1,
  "name": "Jane Smith",
  "dob": "1995-08-15"
}
```

### Delete User

`DELETE /users/:id`

Response:

- `204 No Content`

## Validation Rules

- `name` is required
- `dob` is required and must be in `YYYY-MM-DD` format
- `id` path parameters must be valid integers

## Notes

- Age is calculated dynamically from `dob` when fetching a single user or listing users.
- The handler layer returns validation and parsing errors as JSON responses.
