

**Task Management API with JWT Authentication**

**Introduction**

This document outlines the functionalities, API endpoints, usage instructions, and authentication process for a Task Management API built with Go and Gin framework. It utilizes JWT (JSON Web Token) for user authentication and authorization.


**Running the API Server**

1. Run `go run main.go` to start the server.
2. The API will be accessible by default on port 8080 (http://localhost:8080/).

**API Endpoints**

**User Management:**

* **POST /register**
    * Body: JSON object containing username and password.
    * Response: 201 Created with the newly created user object (excluding password) upon successful registration.
    * Example request:
      ```json
      {
        "username": "testuser",
        "password": "testpassword"
      }
      ```
* **POST /login**
    * Body: JSON object containing username and password for authentication.
    * Response: 200 OK with a JSON object containing the JWT token on successful login.
    * Example request:
      ```json
      {
        "username": "testuser",
        "password": "testpassword"
      }
      ```
    * **Authorization:** None required for login.

**Task Management:**

* **GET /tasks**
    * Retrieves a list of all tasks. Requires a valid JWT token in the Authorization header.
    * Response: 200 OK with a JSON array containing task objects.
    * Example request (with Authorization header):
      ```
      GET /tasks
      Authorization: Bearer <your_jwt_token>
      ```
* **GET /tasks/:id**
    * Retrieves a specific task by its ID. Requires a valid JWT token.
    * Response: 200 OK with a JSON object representing the task or 404 Not Found if the ID is invalid.
    * Example request:
      ```
      GET /tasks/1
      Authorization: Bearer <your_jwt_token>
      ```
* **POST /tasks**
    * Creates a new task. Requires a valid JWT token.
    * Body: JSON object containing the task details.
    * Response: 201 Created with the newly created task object.
    * Example request:
      ```
      POST /tasks
      Content-Type: application/json
      Authorization: Bearer <your_jwt_token>

      {
        "title": "Task title",
        "description": "Task description"
      }
      ```
* **PUT /tasks/:id**
    * Updates a specific task by its ID. Requires a valid JWT token.
    * Body: JSON object containing the updated task details.
    * Response: 200 OK with the updated task object or 404 Not Found if the ID is invalid.
    * Example request:
      ```
      PUT /tasks/1
      Content-Type: application/json
      Authorization: Bearer <your_jwt_token>

      {
        "title": "Updated task title"
      }
      ```
* **DELETE /tasks/:id**
    * Deletes a specific task by its ID. Requires a valid JWT token.
    * Response: 204 No Content on successful deletion or 404 Not Found if the ID is invalid.
    * Example request:
      ```
      DELETE /tasks/1
      Authorization: Bearer <your_jwt_token>
      ```

**Authentication**

The API uses JWTs for authentication and authorization. To access protected routes (like `GET /tasks`), you'll need to obtain a JWT token by logging in with a valid username and password.

1. **Login:** Send a POST request to `/login` with the user's credentials. Upon successful login, the server will respond with a JWT token.
2. **Retrieve Tasks:** Include the JWT token in the Authorization header of subsequent requests to access protected resources:
   ```
   Authorization: Bearer <your_jwt_token>
   ```

**Error Handling**

The API will return standard HTTP status codes to indicate success or failure of requests. Refer to the specific endpoint descriptions for expected response codes

POSTMAN DOCUMENTATION: https://documenter.getpostman.com/view/37289771/2sA3rzKsPp
