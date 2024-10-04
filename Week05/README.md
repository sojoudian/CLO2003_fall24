
# Go CRUD API (Version 1.23)

This is a simple CRUD (Create, Read, Update, Delete) API built with Go version 1.23 using only the built-in Go packages. This API demonstrates the basic principles of a CRUD application by handling HTTP requests for managing a collection of users stored in-memory (simulating a database).

## Folder Structure
In this case, we have the entire application logic in one `main.go` file, which contains all the code necessary for the API.

## Code Breakdown

### 1. Model: `User`
The `User` struct represents a user in our application. It has three fields:
- `ID`: The unique identifier for the user (auto-incremented).
- `Name`: The name of the user.
- `Email`: The user's email address.

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```

### 2. In-memory Storage
We use a slice (`[]User`) to store users in-memory, along with an auto-incrementing `nextID` to assign unique IDs to each user. Since this is an in-memory solution, the data is not persisted and will reset each time the server restarts.

```go
var users = []User{}
var nextID = 1
```

### 3. HTTP Handlers
We created one HTTP handler called `UserHandler`, which handles different HTTP methods such as GET, POST, PUT, and DELETE.

- **GET**: Retrieves all users or a single user by ID.
- **POST**: Creates a new user.
- **PUT**: Updates an existing user by ID.
- **DELETE**: Deletes a user by ID.

### 4. Routing
We use Go's `net/http` package to handle routing. The following routes are defined:
- `GET /users`: Retrieves all users.
- `POST /users`: Creates a new user.
- `GET /users/{id}`: Retrieves a specific user by ID.
- `PUT /users/{id}`: Updates a user by ID.
- `DELETE /users/{id}`: Deletes a user by ID.

### 5. Start the Server
The server is started using Go's `http.ListenAndServe` function, which listens on port `8080`.

```go
func main() {
    // Initialize the HTTP router and routes
    http.HandleFunc("/users", UserHandler)
    http.HandleFunc("/users/", UserHandler)

    // Start the HTTP server
    log.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
```

## Testing the API

You can test the API using `curl` commands or a tool like Postman.

### 1. Get All Users
To retrieve all users, you can use the following `curl` command:
```bash
curl http://localhost:8080/users
```

### 2. Create a User
To create a new user, use this `curl` command:
```bash
curl -X POST -d '{"name": "John Doe", "email": "john@example.com"}' -H "Content-Type: application/json" http://localhost:8080/users
```

### 3. Get a User by ID
To retrieve a user by ID, use the following `curl` command (replace `1` with the desired user ID):
```bash
curl http://localhost:8080/users/1
```

### 4. Update a User
To update a user by ID, use this `curl` command:
```bash
curl -X PUT -d '{"name": "Jane Doe", "email": "jane@example.com"}' -H "Content-Type: application/json" http://localhost:8080/users/1
```

### 5. Delete a User
To delete a user by ID, use this `curl` command:
```bash
curl -X DELETE http://localhost:8080/users/1
```

## Explanation of CRUD Operations
CRUD is an acronym for the four basic operations that can be performed on data in persistent storage:
- **Create**: Adding a new record (HTTP `POST`).
- **Read**: Retrieving records (HTTP `GET`).
- **Update**: Modifying an existing record (HTTP `PUT`).
- **Delete**: Removing a record (HTTP `DELETE`).

Each of these operations corresponds to specific HTTP methods in a RESTful API. In this application, these operations are applied to manage users.

## Requirements

- Go 1.23+
- Curl or Postman for testing the API.

## Running the Application

1. Save the code in `main.go`.
2. Run the application using:
   ```bash
   go run main.go
   ```

The server will start at `http://localhost:8080`, and you can use the API endpoints described above to interact with it.

## Conclusion

This is a basic example of a CRUD API built with Go using only the built-in standard library. It demonstrates how to create and expose HTTP endpoints for managing data in a RESTful way, all within a single file.

