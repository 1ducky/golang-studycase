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

## Running the Application

Ensure you have your `.env` configured properly. The server runs on port `8000` by default.

```bash
go run cmd/api/main.go
```

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
