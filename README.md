# Go REST API

This is a REST API built with Go, featuring dependency injection, custom middleware (logging and JWT authentication), and a modular project structure.

## Features

- JWT Authentication
- CRUD Todo
- Image Upload
- CSV Bulk Import
- Validation
- Transaction
- Logging Middleware
- Worker Pool
- Pipeline Streaming

## Flows
```bash 
Client
   │
HTTP Handler
   │
Mapper
   │
Service
   │
Repository
   │
MySQL/Postgre/etc
```

![Flow Diagram](assets/diagram/Stream_Flow_Diagram.png)


## Configuration

The application is configured using environment variables. You can find the required variables in the `.env.example` file.

Create a `.env` file in the root directory:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=root
DB_PASSWORD=secret
DB_NAME=rest_api
SECRET_AUTH=your_secret_key
```

## Database Setup

Before running the application, make sure you import the initial database schema to your MySQL database. You can find the schema file in the `migrate/mysql/` directory.

To import it, run the following command in your terminal (make sure to replace the username, password, and database name with your own):
```bash
mysql -u root -p rest_api < migrate/mysql/000_initial.sql
```

## Running the Application

Ensure you have your `.env` configured properly. The server runs on port `8000` by default.

```bash
go run cmd/api/main.go
```

## Internal Package Mapper

The `internal` directory contains the core application logic, categorized by domain and technical concerns. Here is the mapping of each package and its corresponding environment variable dependencies:

| Package | Description | Environment Variable (`.env`) |
| :--- | :--- | :--- |
| `application` | Core application setup, wiring, and interfaces. | - |
| `auth` | Authentication logic, including JWT generation and validation. | `SECRET_AUTH` |
| `csv` | Utilities for processing and parsing CSV files. | - |
| `database` | Database connection setup and transaction management. | `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME` |
| `http` | HTTP server configuration, middlewares, and common response handlers. | - |
| `image` | Handling image/file uploads and local storage interactions. | `LOCAL_STORAGE_PATH` |
| `logging` | Custom logging implementations and middleware. | - |
| `pipeline` | Concurrency patterns (e.g., pipeline streaming) implementations. | - |
| `query` | Database query builders or pagination helpers. | - |
| `reader` | Utilities for file or stream reading operations. | - |
| `todos` | Todo domain feature (Handlers, Services, Repositories). | - |
| `user` | User domain feature (Handlers, Services, Repositories). | - |
| `utils` | General purpose helper functions and utilities. | - |

## Bootstrap Process

The application follows a structured initialization flow in `main.go`:
1. **Environment & Configuration**: Loads `.env` file (`config.LoadEnv()`) and initializes the config object.
2. **Database Initialization**: Connects to the database and sets up a DB transaction manager.
3. **Repository Layer**: Instantiates repositories (`TodoLogRepository`, `TodoMemory`, `AuthRepository`).
4. **Service Layer**: Instantiates services with their required dependencies injected (`TodoService`, `JWTService`, `AuthService`).
5. **Handler Layer**: Instantiates HTTP handlers (`TodoHandler`, `AuthHandler`, `ImageHandler`).
6. **Router & Endpoints**: Creates an `http.ServeMux` and registers all routes.
7. **Middlewares**: Chains global middlewares (e.g., Logging and Authentication) wrapping the multiplexer.
8. **Server Start**: Configures the HTTP server (e.g., timeouts, port `:8000`) and starts listening for requests (`server.ListenAndServe()`).

## Endpoints

Based on the available handlers, the API provides the following services and their functionalities:

### Auth Service
- **Login** (`POST /login`): Authenticates a user and returns a JWT token. (Requires JSON payload with credentials)

### Todos Service
- **GetAll** (`GET /todos`): Retrieves all todos. Can be filtered by `id` query parameter.
- **GetById** (`GET /todos?id={id}`): Retrieves a specific todo by `id` query parameter.
- **Create** (`POST /todos`): Creates a new todo. (Requires Auth & JSON payload)
- **Update** (`PATCH /todos`): Updates an existing todo. (Requires Auth & JSON payload)
- **Delete** (`DELETE /todos`): Deletes a todo. (Requires Auth & JSON payload)
- **Upload CSV** (`POST /todos/bulk`): Bulk creates todos from an uploaded CSV file. (Requires Auth & Multipart form data)

### Image Service
- **Upload** (`POST /upload`): Uploads an image file (Multipart form data). Requires Auth.
