
# Simple HTTP Server in Docker

## Overview

This project demonstrates how to create and run a simple HTTP server using Python, packaged in a Docker container. The HTTP server responds to GET requests with a "Hello, world!" message. The setup is lightweight, easy to use, and perfect for learning the basics of Docker and Python HTTP servers.

---

## Table of Contents

1. [Project Structure](#project-structure)
2. [Requirements](#requirements)
3. [Dockerfile Explanation](#dockerfile-explanation)
4. [Python Server Script](#python-server-script)
5. [How to Build and Run](#how-to-build-and-run)
6. [Testing the Server](#testing-the-server)
7. [Key Concepts](#key-concepts)

---

## Project Structure

```
.
├── Dockerfile    # Docker configuration file to create the container
├── server.py     # Python script for the HTTP server
```

---

## Requirements

1. Docker installed on your system. [Download Docker](https://www.docker.com/products/docker-desktop)
2. Basic knowledge of Docker commands and Python.

---

## Dockerfile Explanation

**Dockerfile Contents:**
```dockerfile
FROM python:3.12-slim

WORKDIR /app

COPY server.py .

EXPOSE 5002

CMD [ "python3", "server.py" ]
```

- **Base Image**: Uses `python:3.12-slim`, a minimal Python 3.12 image.
- **Working Directory**: Sets `/app` as the working directory within the container.
- **File Copying**: Copies the `server.py` script into the container's `/app` directory.
- **Port Exposure**: Exposes port `5002` for communication.
- **Default Command**: Runs `server.py` using Python when the container starts.

---

## Python Server Script

The `server.py` script is a basic HTTP server built with Python's `http.server` module.

**Key Features:**
- Handles GET requests and responds with "Hello, world!".
- Runs on port `5002` and listens on all network interfaces (`0.0.0.0`).

**Server Script:**
```python
from http.server import HTTPServer, BaseHTTPRequestHandler

class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/plain")
        self.end_headers()
        self.wfile.write(b"Hello, world!")

def run(server_class=HTTPServer, handler_class=SimpleHTTPRequestHandler):
    server_address = ('0.0.0.0', 5002)
    httpd = server_class(server_address, handler_class)
    print("Serving on http://localhost:5002")
    httpd.serve_forever()

if __name__ == "__main__":
    run()
```

---

## How to Build and Run

1. **Build the Docker Image**:
   ```bash
   docker build -t simple-http-server .
   ```

2. **Run the Docker Container**:
   ```bash
   docker run -p 5002:5002 simple-http-server
   ```

3. **Stop the Container**:
   To stop the running container, press `Ctrl+C` or use:
   ```bash
   docker stop <container-id>
   ```

---

## Testing the Server

- Open a browser and visit: [http://localhost:5002](http://localhost:5002)
- Use `curl` to test:
  ```bash
  curl http://localhost:5002
  ```

**Expected Output**:
```
Hello, world!
```

---

## Key Concepts

1. **Docker Containerization**:
   - Encapsulates the application and its environment for consistency.
   - Lightweight and portable compared to virtual machines.

2. **Python HTTP Server**:
   - Simple and effective for basic web serving tasks.

3. **Port Mapping**:
   - Maps the container's internal port (`5002`) to the host machine's port (`5002`).

---

## Conclusion

This project showcases a basic example of Docker and Python integration. By running the `server.py` script in a container, you gain insights into:
- Dockerfile creation and configuration.
- Building and running containers.
- Testing applications inside Docker.

Feel free to extend this setup for more complex applications or microservices.

---

