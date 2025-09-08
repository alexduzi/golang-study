# Go Testing Guide

This project uses Go's built-in testing framework. Follow the instructions below to run tests effectively.

## Prerequisites

- Go 1.16 or higher installed
- Project dependencies installed (`go mod tidy`)

## Basic Test Commands

### Run All Tests
```bash
# Run all tests in the current directory
go test .

# Run all tests with verbose output
go test . -v

# Run all tests in all subdirectories
go test ./...

# Run all tests in all subdirectories with verbose output
go test ./... -v
```

### Run Specific Tests
```bash
# Run tests matching a pattern
go test -run TestFunctionName

# Run tests in a specific package
go test ./pkg/mypackage

# Run a specific test function with verbose output
go test -run TestSpecificFunction -v
```

## Test Coverage

### Generate Coverage Report
```bash
# Generate coverage profile
go test -coverprofile=coverage.out

# Generate coverage for all packages
go test ./... -coverprofile=coverage.out

# View coverage report in terminal
go tool cover -func=coverage.out

# Generate HTML coverage report
go tool cover -html=coverage.out -o coverage.html
```

### Coverage Options
```bash
# Show coverage percentage while running tests
go test -cover

# Generate coverage with verbose output
go test -coverprofile=coverage.out -v

# Set coverage mode (set, count, atomic)
go test -covermode=count -coverprofile=coverage.out
```

## Benchmark Tests

```bash
# Run benchmark tests
go test -bench=.

# Run benchmarks with memory allocation stats
go test -bench=. -benchmem

# Run benchmarks multiple times for accuracy
go test -bench=. -count=5

# Run specific benchmark
go test -bench=BenchmarkFunctionName
```

## Additional Testing Options

### Parallel Execution
```bash
# Run tests in parallel
go test -parallel 4

# Disable parallel execution
go test -parallel 1
```

### Timeout and Performance
```bash
# Set test timeout
go test -timeout 30s

# Run tests with race condition detection
go test -race

# Show detailed timing information
go test -v -test.v
```

### Short Tests
```bash
# Skip long-running tests (requires testing.Short() in code)
go test -short
```

## Common Test Workflows

### Development Workflow
```bash
# Quick test run during development
go test . -v -short

# Full test suite with coverage
go test ./... -coverprofile=coverage.out -v
go tool cover -html=coverage.out
```

### CI/CD Pipeline
```bash
# Comprehensive test run for CI
go test ./... -race -coverprofile=coverage.out -covermode=atomic -v
go tool cover -func=coverage.out
```

### Performance Testing
```bash
# Run benchmarks and regular tests
go test -bench=. -benchmem -v
```

## Test File Structure

Tests should be placed in files ending with `_test.go` in the same package as the code being tested:

```
project/
├── main.go
├── main_test.go
├── pkg/
│   ├── utils/
│   │   ├── utils.go
│   │   └── utils_test.go
│   └── handlers/
│       ├── handlers.go
│       └── handlers_test.go
└── README.md
```

## Writing Tests

### Basic Test Function
```go
func TestFunctionName(t *testing.T) {
    result := FunctionToTest()
    expected := "expected_value"
    
    if result != expected {
        t.Errorf("Expected %s, got %s", expected, result)
    }
}
```

### Table-Driven Tests
```go
func TestFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"case1", "input1", "output1"},
        {"case2", "input2", "output2"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := FunctionToTest(tt.input)
            if result != tt.expected {
                t.Errorf("Expected %s, got %s", tt.expected, result)
            }
        })
    }
}
```

## Useful Testing Packages

- `testing` - Built-in Go testing package
- `testify/assert` - Popular assertion library
- `testify/mock` - Mocking framework
- `go-sqlmock` - SQL mocking
- `httptest` - HTTP testing utilities

## Example Commands Summary

```bash
# Most commonly used commands:
go test . -v                              # Run tests with output
go test ./... -v                          # Run all tests in project
go test -coverprofile=coverage.out        # Generate coverage
go tool cover -html=coverage.out          # View coverage report
go test -race                             # Test with race detection
go test -bench=.                          # Run benchmarks
```

## Troubleshooting

- **Tests not found**: Ensure test files end with `_test.go`
- **Import issues**: Check that test files are in the same package
- **Coverage not working**: Make sure you're in the correct directory
- **Slow tests**: Use `-short` flag or check for infinite loops

For more information, see the [official Go testing documentation](https://golang.org/pkg/testing/).