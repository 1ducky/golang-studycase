# Go REST API

This is a REST API built with Go, featuring dependency injection, custom middleware (logging and JWT authentication), and a modular project structure.

## Features

- **Authentication**: JWT-based authentication.
- **Database**: Configured via environment variables.
- **Middleware**: Logging and Authentication middlewares.

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
- **Login**: Authenticates a user and returns a JWT token. (Requires JSON payload with credentials)

### Todos Service
- **GetAll**: Retrieves all todos. Can be filtered by `id` query parameter.
- **GetById**: Retrieves a specific todo by `id` query parameter.
- **Create**: Creates a new todo. (Requires Auth & JSON payload)
- **Update**: Updates an existing todo. (Requires Auth & JSON payload)
- **Delete**: Deletes a todo. (Requires Auth & JSON payload)

### Image Service
- **Upload**: Uploads an image file (Multipart form data). Requires Auth.
