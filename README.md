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

### Option 1: Using Cloud PostgreSQL (Aiven, Neon, etc.)

1. Get your PostgreSQL connection URL from your cloud provider (e.g., Aiven, Neon, Supabase, ElephantSQL)

2. Create a `.env` file in the root directory:
```env
DATABASE_URL=postgres://username:password@hostname:port/database?sslmode=require
SERVER_PORT=3000
```

3. Run migration:
```bash
go run cmd/migrate/main.go
```

4. Run the application:
```bash
go run cmd/server/main.go
```

API will be available at `http://localhost:3000`

### Option 2: With Docker (Local)

```bash
docker-compose up --build
```

API will be available at `http://localhost:3000`

### Option 3: Local PostgreSQL

1. Create PostgreSQL database:
```sql
CREATE DATABASE userdb;
```

2. Create `.env` file:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=userdb
SERVER_PORT=3000
```

3. Run migration:
```bash
go run cmd/migrate/main.go
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
| DATABASE_URL | -          | Full PostgreSQL URL (for cloud DBs) |
| DB_HOST      | localhost  | Database host        |
| DB_PORT      | 5432       | Database port        |
| DB_USER      | postgres   | Database user        |
| DB_PASSWORD  | postgres   | Database password    |
| DB_NAME      | userdb     | Database name        |
| SERVER_PORT  | 3000       | Server port          |

> **Note:** If `DATABASE_URL` is set, it takes priority over individual DB variables.

## Testing with Postman

1. Start the server: `go run cmd/server/main.go`
2. Use Base URL: `http://localhost:3000`

| Method | Endpoint | Body (JSON) |
|--------|----------|-------------|
| POST | `/users` | `{"name": "John", "dob": "2000-05-15"}` |
| GET | `/users` | - |
| GET | `/users/:id` | - |
| PUT | `/users/:id` | `{"name": "Jane", "dob": "1999-03-20"}` |
| DELETE | `/users/:id` | - |

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
