
# Testing in Go

---

## Overview

Testing is a critical part of software development. In Go, testing is simple, integrated, and powerful. The `testing` package provides tools to write unit tests, benchmarks, and example-based tests. This README explains how to create and run tests for your Go application.

---

## Why Testing in Go?

1. **Built-in Testing Framework**:
   - Go comes with a native `testing` package, eliminating the need for external libraries.
2. **Simplified Syntax**:
   - Tests are easy to write and execute with minimal setup.
3. **Automated Workflow**:
   - Test execution is integrated with the Go toolchain.

---

## Structure of a Test File

### Naming Conventions

1. **Test Files**:
   - Test files must have the `_test.go` suffix (e.g., `main_test.go`).
2. **Test Functions**:
   - Test functions must start with `Test` and take a single `*testing.T` parameter.

---

## Example: `main_test.go`

This test file validates the functionality of the Go server:

```go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestServer(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    w := httptest.NewRecorder()

    fs := http.FileServer(http.Dir("./static"))
    fs.ServeHTTP(w, req)

    if w.Result().StatusCode != http.StatusOK {
        t.Errorf("Expected status 200, but got %d", w.Result().StatusCode)
    }
}
```

### Key Functions
1. **`httptest.NewRequest`**:
   - Simulates an HTTP request.
2. **`httptest.NewRecorder`**:
   - Captures the server's response for validation.
3. **Assertions**:
   - `t.Errorf` is used to log errors when test cases fail.

---

## Running Tests

1. Execute tests with:
   ```bash
   go test ./...
   ```
   - This command runs all test files in the current and subdirectories.

2. To see detailed output, use:
   ```bash
   go test -v ./...
   ```

---

## Writing Effective Tests

1. **Isolate Units**:
   - Focus on individual functions or components.
2. **Use Test Tables**:
   - Create a slice of test cases to simplify repetitive tests.

Example:
```go
func TestAddition(t *testing.T) {
    testCases := []struct {
        a, b, expected int
    }{
        {1, 2, 3},
        {2, 3, 5},
        {10, -5, 5},
    }

    for _, tc := range testCases {
        result := Add(tc.a, tc.b)
        if result != tc.expected {
            t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
        }
    }
}
```

---

## Advanced Testing Features

1. **Benchmarking**:
   - Measure performance of your code using `func BenchmarkXxx(b *testing.B)`.

Example:
```go
func BenchmarkAddition(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}
```

2. **Example Tests**:
   - Demonstrate code usage in documentation.

Example:
```go
func ExampleAdd() {
    fmt.Println(Add(1, 2))
    // Output: 3
}
```

---

## Best Practices for Testing in Go

1. **Test Coverage**:
   - Aim for high test coverage but focus on critical paths.
   - Use:
     ```bash
     go test -cover ./...
     ```
2. **Avoid Flaky Tests**:
   - Ensure tests are deterministic and repeatable.
3. **CI Integration**:
   - Automate tests with Continuous Integration (CI) pipelines.

---

## Conclusion

Testing in Go ensures the reliability, performance, and maintainability of your applications. By leveraging the `testing` package, you can write effective unit tests, benchmarks, and examples with ease. The provided examples and best practices will help you create a robust testing strategy for your Go projects.



# Test WebServer 
 ```bash
go test -v
ls
go test
go test -v
go test -v -run TestFileServer
go test ./...
```
