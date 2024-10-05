
# Go API Server

This is a simple Go API server that supports both GET and POST requests. The server allows the user to specify the port number via a command-line argument.

## Features

- **GET /**: Returns a welcome message.
- **POST /post**: Accepts JSON data and returns the same data in the response.

## Requirements

- Go 1.16 or later.

## How to Run

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/go-api-server.git
cd go-api-server
```

### 2. Run the server

You can specify the port number when starting the server. If no port number is provided, it defaults to `8091`.

```bash
go run main.go -port=8080
```

This will start the server on `localhost:8080`.

If you want to use the default port (8091), simply run:

```bash
go run main.go
```

### 3. Test the Endpoints

#### GET Request

You can test the GET request by opening the following URL in your browser or using `curl`:

```bash
curl http://localhost:8080/
```

You should see the following response:

```text
Welcome to the Go API
```

#### POST Request

To test the POST endpoint, you can use `curl` to send a JSON payload:

```bash
curl -X POST -d '{"name": "John", "content": "Hello World!"}' http://localhost:8080/post -H "Content-Type: application/json"
```

The response should return the same JSON you sent:

```json
{
    "name": "John",
    "content": "Hello World!"
}
```

## Customization

To change the port, use the `-port` flag when running the server. For example, to run on port 3000:

```bash
go run main.go -port=3000
```

## Project Structure

```bash
.
├── main.go        # Main Go file containing server logic
└── README.md      # Project documentation
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

