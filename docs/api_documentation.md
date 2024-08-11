
# Task Management API Documentation

## Overview

The Task Management API is designed for managing tasks with JWT-based authentication and authorization. The API supports operations for both users and admins, allowing task creation, retrieval, update, and deletion. Authentication is handled via JWT tokens, which are required for accessing protected endpoints.

## Base URL

The base URL for the API is:
- **Development Environment**: `http://localhost:9090`

## Authentication

All endpoints that require authentication need a JWT token in the `Authorization` header. The token should be in the format:
```
Authorization: Bearer <token>
```

Replace `<token>` with the actual JWT obtained during the login process.

## Endpoints

### User Endpoints

#### 1. User Registration
- **Endpoint**: `POST /register`
- **Description**: Registers a new user. Role should be `user`.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string",
    "role": "user"
  }
  ```
- **Response**: Returns the created user object with a unique ID.

#### 2. User Login
- **Endpoint**: `POST /login`
- **Description**: Authenticates a user and returns a JWT token.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Response**: Returns the JWT token.

#### 3. Create Task for User
- **Endpoint**: `POST /tasks`
- **Description**: Creates a new task for the authenticated user.
- **Request Body**:
  ```json
  {
    "title": "string",
    "description": "string",
    "duedate": "ISO8601 datetime",
    "status": "TaskStatus"
  }
  ```
- **Response**: Returns the created task with a unique ID.

#### 4. Get All Tasks for User
- **Endpoint**: `GET /tasks`
- **Description**: Retrieves all tasks associated with the authenticated user.
- **Response**: Returns a list of tasks.

#### 5. Get Task by ID for User
- **Endpoint**: `GET /tasks/{id}`
- **Description**: Retrieves a specific task by its ID for the authenticated user.
- **Response**: Returns the task details.

#### 6. Update Task for User
- **Endpoint**: `PUT /tasks/{id}`
- **Description**: Updates an existing task for the authenticated user.
- **Request Body**:
  ```json
  {
    "title": "string",
    "description": "string",
    "duedate": "ISO8601 datetime",
    "status": "TaskStatus"
  }
  ```
- **Response**: Returns the updated task.

#### 7. Delete Task for User
- **Endpoint**: `DELETE /tasks/{id}`
- **Description**: Deletes a specific task by its ID for the authenticated user.
- **Response**: Returns a confirmation of successful deletion.

### Admin Endpoints

#### 1. Register an Admin
- **Endpoint**: `POST /register`
- **Description**: Registers a new admin user.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string",
    "role": "admin"
  }
  ```
- **Response**: Returns the created admin user object with a unique ID.

#### 2. Admin Login
- **Endpoint**: `POST /login`
- **Description**: Authenticates an admin and returns a JWT token.
- **Request Body**:
  ```json
  {
    "username": "string",
    "password": "string",
    "role": "admin"
  }
  ```
- **Response**: Returns the JWT token.

#### 3. Create Task by Admin
- **Endpoint**: `POST /tasks`
- **Description**: Creates a new task for any user. Admins can specify `userID`.
- **Request Body**:
  ```json
  {
    "title": "string",
    "description": "string",
    "duedate": "ISO8601 datetime",
    "status": "TaskStatus",
    "userID": "integer"
  }
  ```
- **Response**: Returns the created task with a unique ID.

#### 4. Get All Tasks for Admin
- **Endpoint**: `GET /tasks`
- **Description**: Retrieves all tasks in the system, accessible to admin users.
- **Response**: Returns a comprehensive list of all tasks.

#### 5. Get Task by ID for Admin
- **Endpoint**: `GET /tasks/{id}`
- **Description**: Retrieves a specific task by its ID, accessible to admins.
- **Response**: Returns the task details.

#### 6. Update Task by ID for Admin
- **Endpoint**: `PUT /tasks/{id}`
- **Description**: Updates a specific task by its ID. Admins can reassign tasks to different users.
- **Request Body**:
  ```json
  {
    "title": "string",
    "description": "string",
    "duedate": "ISO8601 datetime",
    "status": "TaskStatus",
    "userID": "integer"
  }
  ```
- **Response**: Returns the updated task.

#### 7. Delete Task by ID for Admin
- **Endpoint**: `DELETE /tasks/{id}`
- **Description**: Deletes a specific task by its ID. Admins can delete any task in the system.
- **Response**: Returns a confirmation of successful deletion.

## TaskStatus Enum

The `status` field can have one of the following values:
- `complete`
- `in_progress`
- `started`

## Error Handling

The API may return the following HTTP status codes:
- **400 Bad Request**: Invalid request payload or parameters.
- **401 Unauthorized**: Invalid or missing authentication token.
- **403 Forbidden**: Insufficient permissions to access or modify the resource.
- **404 Not Found**: Resource (user or task) not found.
- **500 Internal Server Error**: Unexpected server error.

## Example Requests

### User Registration
```http
POST /register
Content-Type: application/json

{
  "username": "john_doe",
  "password": "securepassword",
  "role": "user"
}
```

### User Login
```http
POST /login
Content-Type: application/json

{
  "username": "john_doe",
  "password": "securepassword"
}
```

### Create Task for User
```http
POST /tasks
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "New Task",
  "description": "Task description",
  "duedate": "2024-08-15T14:30:00Z",
  "status": "in_progress"
}
```

### Update Task for User
```http
PUT /tasks/{id}
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Updated Task Title",
  "description": "Updated description",
  "duedate": "2024-08-15T14:30:00Z",
  "status": "complete"
}
```

### Get All Tasks for User
```http
GET /tasks
Authorization: Bearer <token>
```

### Delete Task for User
```http
DELETE /tasks/{id}
Authorization: Bearer <token>
```

## Conclusion

This API provides a robust set of endpoints for managing tasks with clear role-based access control. Ensure to include JWT tokens for authenticated endpoints and follow the request and response formats as described.

**Error Handling**

The API will return standard HTTP status codes to indicate success or failure of requests. 

API DOCUMENTATION: [https://documenter.getpostman.com/view/37289771/2sA3rzKsPp]
