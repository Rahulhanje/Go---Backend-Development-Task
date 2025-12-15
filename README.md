# Go User API

A REST API built with GoFiber to manage users with name and date of birth, calculating age dynamically.

## Tech Stack

- **GoFiber** - Fast HTTP framework
- **PostgreSQL** - Database
- **SQLC** - Type-safe SQL
- **Uber Zap** - Structured logging
- **go-playground/validator** - Request validation

## Project Structure

```
├── cmd/server/          # Application entry point
├── config/              # Configuration & database connection
├── db/
│   ├── migrations/      # SQL migrations
│   └── sqlc/            # SQLC generated code
└── internal/
    ├── handler/         # HTTP handlers
    ├── service/         # Business logic (age calculation)
    ├── repository/      # Database operations
    ├── routes/          # Route setup
    ├── middleware/      # Request ID & logging
    ├── models/          # Request/Response structs
    └── logger/          # Zap logger setup
```

## API Endpoints

| Method | Endpoint     | Description         |
|--------|-------------|---------------------|
| POST   | /users      | Create a user       |
| GET    | /users/:id  | Get user by ID      |
| PUT    | /users/:id  | Update user         |
| DELETE | /users/:id  | Delete user         |
| GET    | /users      | List all users      |

## Request/Response Examples

### Create User
```json
POST /users
{
  "name": "John Doe",
  "dob": "2000-05-15"
}

Response:
{
  "id": 1,
  "name": "John Doe",
  "dob": "2000-05-15",
  "age": 25
}
```

## Running the Application

### With Docker (Recommended)

```bash
docker-compose up --build
```

API will be available at `http://localhost:3000`

### Without Docker

1. Create PostgreSQL database:
```sql
CREATE DATABASE userdb;
```

2. Run migration:
```bash
psql -d userdb -f db/migrations/001_create_users_table.up.sql
```

3. Set environment variables:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=userdb
export SERVER_PORT=3000
```

4. Run the application:
```bash
go run cmd/server/main.go
```

## Running Tests

```bash
go test ./... -v
```

## Environment Variables

| Variable      | Default    | Description          |
|--------------|------------|----------------------|
| DB_HOST      | localhost  | Database host        |
| DB_PORT      | 5432       | Database port        |
| DB_USER      | postgres   | Database user        |
| DB_PASSWORD  | postgres   | Database password    |
| DB_NAME      | userdb     | Database name        |
| SERVER_PORT  | 3000       | Server port          |

## Features

- ✅ CRUD operations for users
- ✅ Dynamic age calculation with birthday correction
- ✅ Request validation
- ✅ Structured logging with Zap
- ✅ Request ID middleware
- ✅ Request duration logging
- ✅ Proper error handling with JSON responses
- ✅ Docker support
- ✅ Unit tests for age calculation
