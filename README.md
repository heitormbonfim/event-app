# Event App

An application to save and manage events.

## Technologies

- **Go**: Programming language used for the application.
- **Gin Web Framework**: Handles routing and middleware.
- **SQLite**: Lightweight database stored in `data.db`.
- **JWT**: Used for authentication.
- **Swagger**: Generates API documentation.

## Setup

1. **Clone the repository:**
   ```
   git clone <repository-url>
   ```
2. **Create a `.env` file** in the root directory with the following content:
   ```
   PORT=8080
   JWT_SECRET=your-secret-key
   ```
   Replace `your-secret-key` with a secure secret key for JWT authentication.
3. **Install dependencies:**
   ```
   go mod tidy
   ```
4. **Run database migrations:**
   ```
   go run cmd/migrate/main.go
   ```
   This applies the SQL migration files located in `cmd/migrate/migrations/` to set up the database schema.
5. **Generate Swagger documentation:**
   ```
   swag init --dir cmd/api --parseDependency --parseInternal --parseDepth 1
   ```
   Then, replace the `SwaggerInfo` struct in the generated `docs.go` file with:
   ```go
   var SwaggerInfo = &swag.Spec{
       Version: "1.0",
       Host: "",
       BasePath: "",
       Schemes: []string{},
       Title: "Event APP",
       Description: "An app to save events",
       InfoInstanceName: "swagger",
       SwaggerTemplate: docTemplate,
   }
   ```
   Run the command again to finalize:
   ```
   swag init --dir cmd/api
   ```
6. **Run the application:**
   ```
   go run cmd/api/main.go
   ```
   Alternatively, if you have [Air](https://github.com/air-verse/air) installed for live reloading:
   ```
   air
   ```

## Project Structure

- **`cmd/api/`**: Main application code, including routes, middleware, authentication, and server setup.
- **`cmd/migrate/`**: Database migration scripts and the migration tool.
- **`internal/database/`**: Contains database models and queries (e.g., users, events, attendees).
- **`internal/env/`**: Handles environment variable loading.
- **`docs/`**: Stores generated Swagger documentation files.

## API Documentation

After generating Swagger docs, access the API documentation at `http://localhost:8080/swagger`.
