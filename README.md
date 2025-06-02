# Go User API

A simple REST API built with Go Fiber framework

## Features

- Create user
- User login
- Get all users
- Get user by ID

## Prerequisites

Before you begin, ensure you have the following installed:
- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/)
- [Postman](https://www.postman.com/) (for API testing)

## Installation and Setup

### 1. Clone the Repository

```bash
git clone https://github.com/KulpithaC/go-test.git
cd go-test
```

### 2. Start the Application

Start both the database and application using Docker Compose:

```bash
docker-compose up --build
```

This will:
- Start a PostgreSQL database container
- Build and start the Go application container
- Expose the API on `http://localhost:8080`

### 3. Verify Installation

The application should be running on `http://localhost:8080`. You can verify by checking the logs or making a test request.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/users` | Get all users |
| GET | `/users/:id` | Get user by ID |
| POST | `/users` | Create a user |
| POST | `/login` | User login |

## Testing with Postman

### 1. Import Postman Collection

1. Open Postman
2. Click "Import" 
3. Select the `go-test.postman_collection.json` file from the project root
4. The collection "go-test" will be added to your Postman workspace

### 2. API Usage Examples

#### Create User
```http
POST http://localhost:8080/users
Content-Type: application/json

[{
    "name": "Test",
    "email": "test@mail.com",
    "password": "123456"
}]
```

#### Login
```http
POST http://localhost:8080/login
Content-Type: application/json

{
    "email": "test@mail.com",
    "password": "123456"
}
```

#### Get All Users
```http
GET http://localhost:8080/users
```

#### Get User by ID
```http
GET http://localhost:8080/users/:id
```

### 3. Testing Flow

1. **Create a user** using the "CreateUser" request
2. **Login** with the created user credentials using the "Login" request
3. **Get all users** using the "GetAllUser" request
4. **Get specific user** using the "GetUser" request with a valid user ID

## Project Structure

```
go-test/
├── database/          # Database connection and configuration
├── handlers/          # HTTP request handlers
├── models/           # Data models/structs
├── repository/       # Data access layer
├── services/         # Business logic layer
├── docker-compose.yaml
├── Dockerfile
├── go.mod            # Go module file
├── go.sum            # Go dependencies checksum
├── go-test.postman_collection.json
├── main.go
└── README.md
```

## Troubleshooting

### Common Issues

1. **Port already in use**
   ```bash
   # Check what's using port 8080
   lsof -i :8080
   # Kill the process or change the port in docker-compose.yaml
   ```

2. **Database connection failed**
   ```bash
   # Restart the database container
   docker-compose restart db
   ```

3. **Build errors**
   ```bash
   # Clean Docker cache and rebuild
   docker-compose down
   docker system prune -f
   docker-compose up --build
   ```

### Stop the Application
```bash
docker-compose down
```