package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
)

// User defines a user in our system
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// In-memory storage (simulating a database)
var users = []User{}
var nextID = 1

// UserHandler handles all the CRUD operations for users
func UserHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Extract the ID from the URL (if present)
    idStr := strings.TrimPrefix(r.URL.Path, "/users/")
    id, _ := strconv.Atoi(idStr)

    switch r.Method {
    case http.MethodGet:
        if id > 0 {
            getUserByID(w, r, id)
        } else {
            getAllUsers(w, r)
        }
    case http.MethodPost:
        createUser(w, r)
    case http.MethodPut:
        updateUser(w, r, id)
    case http.MethodDelete:
        deleteUser(w, r, id)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

// Get all users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(users)
}

// Get user by ID
func getUserByID(w http.ResponseWriter, r *http.Request, id int) {
    user, found := getUser(id)
    if !found {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(user)
}

// Create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    user.ID = nextID
    nextID++
    users = append(users, user)
    json.NewEncoder(w).Encode(user)
}

// Update an existing user
func updateUser(w http.ResponseWriter, r *http.Request, id int) {
    var updatedUser User
    if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    updatedUser.ID = id
    if updateExistingUser(&updatedUser) {
        json.NewEncoder(w).Encode(updatedUser)
    } else {
        http.Error(w, "User not found", http.StatusNotFound)
    }
}

// Delete a user by ID
func deleteUser(w http.ResponseWriter, r *http.Request, id int) {
    if deleteExistingUser(id) {
        w.WriteHeader(http.StatusNoContent)
    } else {
        http.Error(w, "User not found", http.StatusNotFound)
    }
}

// In-memory functions (simulating database actions)

// GetUser returns a user by ID
func getUser(id int) (User, bool) {
    for _, user := range users {
        if user.ID == id {
            return user, true
        }
    }
    return User{}, false
}

// UpdateUser updates an existing user
func updateExistingUser(updatedUser *User) bool {
    for i, user := range users {
        if user.ID == updatedUser.ID {
            users[i] = *updatedUser
            return true
        }
    }
    return false
}

// DeleteUser deletes a user by ID
func deleteExistingUser(id int) bool {
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            return true
        }
    }
    return false
}

func main() {
    // Initialize the HTTP router and routes
    http.HandleFunc("/users", UserHandler)     // Handles all user routes
    http.HandleFunc("/users/", UserHandler)    // Handles specific user actions

    // Start the HTTP server
    log.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

