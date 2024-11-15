# Go Testing Example

This repository provides a simple Go application designed to teach the fundamentals of testing in Go. It includes examples of unit tests and benchmarks for basic arithmetic operations.

## Project Structure

- `main.go`: Contains the main application code with basic arithmetic functions (`Add`, `Subtract`, `Multiply`) and a simple demonstration in the `main` function.
- `main_test.go`: Contains the unit tests and benchmarks for the functions defined in `main.go`.
- `go.mod`: Defines the module and Go version.

## Requirements

- Go version 1.23 or later.

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/sojoudian/CLO2003_fall24.git
   cd CLO2003_fall24/Week11
   ```

2. Run the application:

   ```bash
   go run main.go
   ```

   Expected output:
   ```
   Addition:  2025
   Subtraction:  0
   Multiplication:  16
   ```

## Running Tests

To run the tests, use the `go test` command:

```bash
go test ./...
```

Sample output:
```
ok      wk11    0.003s
```

### Benchmarking

To run the benchmarks, use the following command:

```bash
go test -bench=.
```

Sample output:
```
goos: darwin
goarch: amd64
pkg: wk11
BenchmarkAdd-8           260000000            4.58 ns/op
BenchmarkSubtract-8      260000000            4.59 ns/op
PASS
ok      wk11    1.248s
```

## Code Explanation

### `main.go`

This file contains three functions:

- `Add(a, b int) int`: Returns the sum of `a` and `b`.
- `Subtract(a, b int) int`: Returns the result of `a - b`.
- `Multiply(a, b int) int`: Returns the product of `a` and `b`.

### `main_test.go`

This file provides:

1. **Unit Tests**:
   - `TestAdd`: Verifies the correctness of the `Add` function.
   - `TestSubtract`: Verifies the correctness of the `Subtract` function.
   - `TestMultiply`: Verifies the correctness of the `Multiply` function.

2. **Benchmarks**:
   - `BenchmarkAdd`: Measures the performance of the `Add` function.
   - `BenchmarkSubtract`: Measures the performance of the `Subtract` function.

### Sample Test Code

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Fatalf("Add(2, 3) will result %d, but we got %d", result, expected)
    }
}
```

This test checks if the `Add` function correctly computes the sum of `2` and `3`.

### Sample Benchmark Code

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(100, 50)
    }
}
```

This benchmark measures the performance of the `Add` function when adding `100` and `50`.

## Learning Objectives

- Understand how to write and run unit tests in Go.
- Learn how to use the `testing` package for benchmarks.
- Gain familiarity with interpreting test and benchmark results.

## License

This project is open-source and available under the MIT License.

---
Happy Coding!
Maziar
