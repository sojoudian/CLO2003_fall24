
# Full-Stack Go Project with External API, MongoDB, and Docker Integration

---

## Overview

This project integrates the following concepts:
1. **JavaScript Integration**: Fetching data from an external API in `index.html`.
2. **Backend API**: Posting the fetched data (IP address) to a Go backend.
3. **Improved Go Code**: Enhancing the backend for scalability and maintainability.
4. **MongoDB**: Using MongoDB as a Docker container and connecting it with the Go application.
5. **Security**: Securing MongoDB and the Go application.
6. **Multistage Docker Image**: Creating an optimized Docker image for the project.

---

## 1. Using External API in `index.html`

### JavaScript Code:
In `index.html`, we fetch the client's IP address from [ipify API](https://www.ipify.org/) and send it to the backend:

```javascript
fetch('https://api.ipify.org?format=json')
    .then(response => response.json())
    .then(data => {
        fetch('http://localhost/api/visitors', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ ip: data.ip }),
        });
    })
    .catch(error => console.error('Error:', error));
```

---

## 2. Implementing POST API in Go

The Go backend provides a POST endpoint to accept IP addresses and store them in MongoDB.

### Example Code:
```go
package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func initMongoDB() {
    var err error
    client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err = client.Connect(ctx); err != nil {
        log.Fatal(err)
    }
}

func handleVisitors(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
        return
    }
    var visitor struct {
        IP string `json:"ip"`
    }
    if err := json.NewDecoder(r.Body).Decode(&visitor); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := client.Database("mydb").Collection("visitors")
    _, err := collection.InsertOne(context.TODO(), bson.M{"ip": visitor.IP, "timestamp": time.Now()})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func main() {
    initMongoDB()
    http.HandleFunc("/api/visitors", handleVisitors)
    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

## 3. Using MongoDB as a Docker Container

### Command to Run MongoDB:
```bash
docker run -d --name mongodb -p 27017:27017 -v mongodb:/data/db mongo:latest
```

### Explanation of Options:
- `-d`: Run the container in detached mode.
- `--name mongodb`: Assigns the name "mongodb" to the container.
- `-p 27017:27017`: Maps port `27017` on the host to port `27017` on the container.
- `-v mongodb:/data/db`: Persists MongoDB data using a Docker volume.
- `mongo:latest`: Specifies the latest MongoDB image from Docker Hub.

---

## 4. Using MongoDB in Go

### Install MongoDB Go Driver:
```bash
go get go.mongodb.org/mongo-driver/mongo
```

### Connecting to MongoDB:
- Use `mongo.NewClient` and `mongo.Connect` for database connection.

### Example Operations:
- **Insert Data**:
  ```go
  collection.InsertOne(context.TODO(), bson.M{"field": "value"})
  ```
- **Query Data**:
  ```go
  collection.Find(context.TODO(), bson.M{"field": "value"})
  ```

---

## 5. Security Best Practices

### MongoDB Security:
- **Authentication**:
  Enable authentication by creating users with specific roles.
- **TLS Encryption**:
  Use TLS to encrypt communication.

### Go Application Security:
- Validate all inputs to prevent injection attacks.
- Use environment variables for sensitive data like database URIs.

---

## 6. Multistage Dockerfile

### Example Dockerfile:
```dockerfile
# Stage 1: Build Stage
FROM golang:1.20 as builder
WORKDIR /app

# Copy dependencies and source code
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
RUN go build -o server main.go

# Stage 2: Production Stage
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary and static files
COPY --from=builder /app/server .
COPY ./static ./static

# Install CA certificates for secure HTTPS requests
RUN apk --no-cache add ca-certificates

# Expose port and define the command
EXPOSE 8080
CMD ["./server"]
```

---

## Steps to Build and Run the Docker Image

1. **Build the Image**:
   ```bash
   docker build -t go-mongo-app .
   ```

2. **Run the Container**:
   ```bash
   docker run -p 8080:8080 go-mongo-app
   ```

---

## Conclusion

This project demonstrates a full-stack integration of JavaScript, a Go backend, MongoDB, and Docker. The setup ensures scalability, security, and portability, making it production-ready.





# External API and save result to MongoDB

```bash
git clone https://github.com/sojoudian/WINP2000_M06.git
curl https://api.ipify.org\?format\=json
docker run -d --name mongodb -p 27017:27017 -v mongodb:/data/db mongo:latest

go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options

"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
```
## How to test the project
 ```bash
go test -v
ls
go test
go test -v
go test -v -run TestFileServer
go test ./...
```
