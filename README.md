
# Task Manager API

Overview
The Task Management API is a RESTful API that allows users and administrators to manage tasks with JWT-based authentication and authorization. This API supports operations for creating, reading, updating, and deleting tasks. It is designed for managing tasks with role-based access control where users and admins have different levels of access.


## Project Structure

```
task-manager/
├── config/
│   └── config.go
├── controllers/
│   ├── user_controller.go
│   ├── task_controller.go
├── data/
│   ├── user_services.go
│   ├── task_services.go
│  
├── middleware/
│   ├── auth_middleware.go
│   └── admin_middleware.go
├── models/
│   ├── task.go
│   └── user.go
├── router/
│   └── router.go
└── main.go
```

## Features

- **User Registration and Login**: Users can register and log in to receive a JWT token.
- **JWT Authentication**: Secure endpoints using JWT tokens.
- **Role-Based Access Control**: Separate routes and permissions for regular users and admins.
- **Task Management**: Create, read, update, and delete tasks.
- **Task Ownership Validation**: Users can only manage their own tasks, while admins can manage all tasks.

## Installation

1. **Clone the repository**:

    ```sh
    git clone [https://github.com/Simret101/JWT_Authentiction_Auth_TaskManagement]
    cd task-manager
    ```

2. **Install dependencies**:

    Make sure you have Go installed. You can install dependencies by running:

    ```sh
    go mod tidy
    ```

3. **Run the application**:

    ```sh
    go run main.go
    ```

    The application will start on `http://localhost:9090`.


## API Endpoints

### Public Endpoints

- **POST /register**: Register a new user.
- **POST /login**: Login and receive a JWT token.

### User Endpoints (Requires Authentication)

- **GET /tasks**: Retrieve all tasks for the logged-in user.
- **GET /tasks/:id**: Retrieve a specific task by ID (only if the user owns it).
- **POST /tasks**: Create a new task.
- **PUT /tasks/:id**: Update a task by ID (only if the user owns it).
- **DELETE /tasks/:id**: Delete a task by ID (only if the user owns it).

### Admin Endpoints (Requires Admin Role)

- **GET /admin/tasks**: Retrieve all tasks.
- **GET /admin/tasks/:id**: Retrieve a specific task by ID.
- **POST /admin/tasks**: Create a new task.
- **PUT /admin/tasks/:id**: Update any task by ID.
- **DELETE /admin/tasks/:id**: Delete any task by ID.

## Middleware

- **AuthMiddleware**: Ensures that requests include a valid JWT token.
- **AdminMiddleware**: Restricts access to admin-specific routes.
- **TaskOwnershipMiddleware**: Ensures that a user can only manage their own tasks.

## Models

- **User**: Represents a user with `ID`, `Username`, `Password`, and `Role`.
- **Task**: Represents a task with `ID`, `Title`, `Description`, `DueDate`, and `Status`.

## Testing

You can test the API using tools like Postman or curl. Here are some basic steps:

1. **Register a user** via `POST /register`.
2. **Login** to receive a JWT token via `POST /login`.
3. **Create and manage tasks** using the user and admin endpoints.

Make sure to include the JWT token in the `Authorization` header for protected routes.



## Contributing

Contributions are welcome! Please fork this repository and submit a pull request for review.

## Contact

If you have any questions or suggestions, feel free to contact me at /semretb74@gmai.com.

