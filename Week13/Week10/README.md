
# Simple Web Server in Go with Docker Integration

---

## Purpose of the Project
The purpose of this project is to teach the following concepts:
1. Building a simple web server in Go.
2. Serving files over the web using a file server.
3. Creating a basic `index.html` for the web interface.
4. Dockerizing a simple Go application.
5. Creating a Docker image, running it locally, tagging the image, and pushing it to Docker Hub.
6. Using the Docker image anywhere by pulling it from Docker Hub.

---

## Project Structure
- **`main.go`**: Contains the Go code for the web server.
- **`index.html`**: A basic HTML webpage that will be served by the Go server.
- **`Dockerfile`**: Dockerfile for containerizing the Go application.
- **`go.mod`**: Dependency management file for Go modules.

---

## Step-by-Step Guide

### 1. Building the Web Server in Go
The web server is written in Go and serves files from a directory. The main components include:
- `http.FileServer`: Serves static files from a directory.
- `http.Handle`: Maps a path to the file server.
- `http.ListenAndServe`: Starts the server.

#### Example Code (main.go):
```go
package main

import (
    "net/http"
    "log"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

### 2. Creating the `index.html` File
The `index.html` file serves as the primary web page. It includes:
- A responsive design for mobile and desktop viewing.
- Sections for "About Me," "Projects," and "Contact."

Place the `index.html` file in a directory named `static/` to be served by the Go server.

---

### 3. Running the Go Server Locally
1. Install Go on your system.
2. Place the `index.html` in a `static/` folder.
3. Run the following command in the terminal:
   ```bash
   go run main.go
   ```
4. Open `http://localhost:8080` in your browser to see the `index.html` file.

---

### 4. Dockerizing the Go Application
The Dockerfile contains the instructions to build and run the Go application in a container.

#### Dockerfile:
```dockerfile
# Use official Golang image for building the app
FROM golang:1.20 as builder
WORKDIR /app

# Copy Go modules and app code
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the Go binary
RUN go build -o server main.go

# Use a lightweight base image for the runtime
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
COPY ./static ./static

# Expose port 8080 and run the app
EXPOSE 8080
CMD ["./server"]
```

#### Steps to Build and Run the Docker Image
1. Build the Docker image:
   ```bash
   docker build -t my-go-app .
   ```
2. Run the image locally:
   ```bash
   docker run -p 8080:8080 my-go-app
   ```

---

### 5. Pushing the Docker Image to Docker Hub
#### Prerequisites
- A Docker Hub account.
- Login to Docker Hub:
  ```bash
  docker login
  ```

#### Steps
1. Tag the Docker image:
   ```bash
   docker tag my-go-app <your-dockerhub-username>/my-go-app
   ```
2. Push the image to Docker Hub:
   ```bash
   docker push <your-dockerhub-username>/my-go-app
   ```

---

### 6. Using the Docker Image Anywhere
You can pull the Docker image on any system and run it without needing the source code.

#### Steps:
1. Pull the Docker image:
   ```bash
   docker pull <your-dockerhub-username>/my-go-app
   ```
2. Run the image:
   ```bash
   docker run -p 8080:8080 <your-dockerhub-username>/my-go-app
   ```

---

## Conclusion
This project covers the fundamentals of building a web server in Go, creating a Dockerized application, and deploying it using Docker Hub. The steps provide a practical understanding of containerization and web server deployment.




# Understanding Multistage Dockerfile with Example

---

## What is a Multistage Dockerfile?

A multistage Dockerfile is a powerful feature in Docker that allows you to use multiple `FROM` instructions within a single Dockerfile. This is particularly useful for optimizing image sizes, as it enables you to:
1. Build your application in one stage.
2. Copy only the necessary files/binaries into the final, lightweight stage.

The result is a production-ready Docker image that is smaller and more secure.

---

## Advantages of Multistage Builds
1. **Reduced Image Size**:
   - Development tools, temporary files, and unnecessary dependencies are left behind in the build stages.
2. **Simplified Workflow**:
   - No need for separate scripts or manual steps to build and copy artifacts.
3. **Improved Security**:
   - Minimizing the number of layers reduces the attack surface of the container.

---

## Multistage Dockerfile Example

Based on the uploaded Go files, here's an example of a multistage Dockerfile:

```dockerfile
# Stage 1: Build Stage
FROM golang:1.20 as builder
WORKDIR /app

# Copy Go modules and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the application
COPY . .
RUN go build -o server main.go

# Stage 2: Production Stage
FROM alpine:latest
WORKDIR /root/

# Copy the compiled binary and static files from the builder stage
COPY --from=builder /app/server .
COPY ./static ./static

# Expose the necessary port and define the default command
EXPOSE 8080
CMD ["./server"]
```

---

## Step-by-Step Explanation of the Multistage Dockerfile

### 1. **Stage 1: Build Stage**
- **`FROM golang:1.20 as builder`**:
  - The first stage uses the official Golang image for building the application.
  - This stage is named `builder` to refer to it later.

- **`WORKDIR /app`**:
  - Sets the working directory for subsequent commands.

- **Copy Go Modules**:
  ```dockerfile
  COPY go.mod go.sum ./
  RUN go mod download
  ```
  - This ensures dependencies are downloaded efficiently using the Go module system.

- **Copy Source Code**:
  ```dockerfile
  COPY . .
  RUN go build -o server main.go
  ```
  - Copies the source code into the container and builds the Go application into an executable binary named `server`.

---

### 2. **Stage 2: Production Stage**
- **`FROM alpine:latest`**:
  - This stage uses a minimal base image (`alpine`) to create a lightweight production image.

- **`COPY --from=builder`**:
  ```dockerfile
  COPY --from=builder /app/server .
  COPY ./static ./static
  ```
  - Only the compiled binary (`server`) and required files (e.g., `static/`) are copied from the `builder` stage.

- **Expose Port and Define Command**:
  ```dockerfile
  EXPOSE 8080
  CMD ["./server"]
  ```
  - The `EXPOSE` command specifies the port for the application.
  - The `CMD` command runs the `server` binary by default.

---

## Why Use Multistage Builds for This Project?

- **Optimized Size**:
  - The final production image only contains the compiled binary and static files, without Go tools or intermediate files.

- **Better Portability**:
  - The lightweight production image is ideal for deployment in resource-constrained environments.

- **Easier Maintenance**:
  - All steps (build and deployment) are consolidated in a single Dockerfile.

---

## Steps to Build and Run the Multistage Docker Image

1. Build the image:
   ```bash
   docker build -t my-multistage-go-app .
   ```
2. Run the container:
   ```bash
   docker run -p 8080:8080 my-multistage-go-app
   ```

---

## Conclusion

Multistage Dockerfiles simplify the process of creating optimized, secure, and production-ready container images. By separating the build and runtime environments, you achieve better performance and maintainability for your applications.



# Docker commands
 ```bash
 docker build -t dw10:0.1 .
 docker run dw10:0.1
 docker run -p 80:80 dw10:0.1
 docker run -p 8033:80 dw10:0.1
 docker login
 docker tag dw10:0.1 maziar/dw10:0.1
 docker push maziar/dw10:0.1
```
