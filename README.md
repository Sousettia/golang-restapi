# Go User Authentication API

A robust REST API built with Go and Gin framework that provides user authentication, management, and real-time monitoring capabilities with MongoDB integration.

## Features

- **User Authentication**: Secure registration and login with JWT tokens
- **Password Security**: Bcrypt hashing for password protection
- **User Management**: Complete CRUD operations for user data
- **JWT Middleware**: Protected routes with token validation
- **MongoDB Integration**: Persistent data storage with MongoDB
- **Background Tasks**: Concurrent user count logging every 10 seconds
- **Request Logging**: Comprehensive logging middleware

## Tech Stack

- **Backend**: Go with Gin framework
- **Database**: MongoDB
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt
- **Environment**: dotenv for configuration

## Project Structure

```
├── config/
│   └── db.go              # MongoDB connection configuration
├── handlers/
│   ├── auth.go            # Registration and login handlers
│   └── user.go            # User CRUD operations
├── middlewares/
│   ├── jwt.go             # JWT authentication middleware
│   └── log.go             # Request logging middleware
├── models/
│   └── user.go            # User data model
├── tasks/
│   └── background.go      # Background task for user count logging
├── utils/
│   └── jwt.go             # JWT utility functions
├── .env                   # Environment variables
└── main.go               # Application entry point
```

## API Endpoints

### Authentication

#### Register User
```http
POST /register
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "1234"
}
```

#### Login User
```http
POST /login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "1234"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUz..."
}
```

### User Management (Protected Routes)

All user management endpoints require authentication. Include the JWT token in the Authorization header
```http
Authorization: Bearer <your-jwt-token>
```

#### Get All Users
```http
GET /api/users
```

#### Get User by ID
```http
GET /api/users/:id
```

#### Update User
```http
PUT /api/users/:id
Content-Type: application/json

{
  "name": "Updated Name",
  "email": "updated@example.com"
}
```

#### Delete User
```http
DELETE /api/users/:id
```

## User Model

The User model includes the following fields:

```go
type User struct {
    ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
    Name      string             `json:"name" bson:"name"`
    Email     string             `json:"email" bson:"email"`
    Password  string             `json:"-" bson:"password"`
    CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
```

## Security Features

- **Password Hashing**: All passwords are hashed using bcrypt before storage
- **JWT Authentication**: Secure token-based authentication with HS256 signing
- **Protected Routes**: User management endpoints are protected by JWT middleware
- **Environment Variables**: Sensitive data like JWT secrets are stored in environment variables

## Background Tasks

The application includes a concurrent background task that logs the current user count every 10 seconds. 
This demonstrates Go's concurrency capabilities and provides real-time monitoring of user registration activity.

## Middleware
### JWT Middleware
- Validates JWT tokens on protected routes
- Extracts user information from tokens
- Blocks unauthorized requests

### Logging Middleware
- Logs all incoming requests
- Includes request method, path, and response time
- Helps with debugging and monitoring

## Testing

**Mainly use Postman for API Testing**