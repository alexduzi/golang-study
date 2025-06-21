# Golang Questions & Answers

A comprehensive collection of 50 Go questions with concise answers, ranging from beginner to expert level.

## 📝 Beginner Level (Questions 1-15)

### 1. What is Go and what are its main characteristics?

**Answer:** Go is a statically typed, compiled language by Google. Key features: simple syntax, fast compilation, built-in concurrency (goroutines), garbage collection, cross-platform compilation.

### 2. What are the basic data types in Go?

**Answer:** 
- Numeric: `int`, `int32`, `int64`, `uint`, `float32`, `float64`
- Boolean: `bool`
- String: `string`
- Aliases: `byte` (uint8), `rune` (int32)

### 3. How do you declare variables in Go?

**Answer:**
```go
var name string = "John"    // explicit type
var age = 25               // type inference
city := "NYC"              // short declaration (functions only)
var x, y int = 10, 20      // multiple variables
```

### 4. What is the difference between `var` and `:=`?

**Answer:**
- `var`: Can be used anywhere, allows explicit typing, creates zero values
- `:=`: Function scope only, type inference, must initialize

### 5. Explain Go's zero values.

**Answer:** Default values when declared without initialization:
- Numbers: `0`, Booleans: `false`, Strings: `""`, Pointers/slices/maps/channels: `nil`

### 6. What are slices and how do they differ from arrays?

**Answer:**
```go
arr := [5]int{1, 2, 3, 4, 5}  // Array: fixed size, value type
slice := []int{1, 2, 3, 4, 5} // Slice: dynamic, reference type
slice = append(slice, 6)       // Dynamic growth
```

### 7. How do you create and use maps in Go?

**Answer:**
```go
m := make(map[string]int)
m["key"] = 42
value, exists := m["key"]  // comma ok idiom
delete(m, "key")
```

### 8. What are pointers in Go?

**Answer:** Variables that store memory addresses. Use `&` to get address, `*` to dereference.
```go
x := 42
p := &x      // p points to x
*p = 100     // modify x through pointer
```

### 9. How do you define and use functions in Go?

**Answer:**
```go
func add(a, b int) int {
    return a + b
}

func divide(a, b int) (int, error) {  // multiple returns
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### 10. What are structs in Go?

**Answer:** Custom types that group related fields together.
```go
type Person struct {
    Name string
    Age  int
}
p := Person{Name: "John", Age: 30}
```

### 11. How do you handle errors in Go?

**Answer:**
```go
result, err := someFunction()
if err != nil {
    return err  // or handle error
}
// use result
```

### 12. What are methods in Go?

**Answer:**
```go
func (p Person) String() string {
    return p.Name
}

func (p *Person) SetAge(age int) {  // pointer receiver
    p.Age = age
}
```

### 13. What is the difference between value and pointer receivers?

**Answer:**
- **Value receiver**: Receives copy, can't modify original
- **Pointer receiver**: Receives pointer, can modify original, more efficient for large structs

### 14. How do you use `range` in Go?

**Answer:**
```go
for index, value := range slice {}
for key, value := range map {}
for _, value := range slice {}    // ignore index
for value := range channel {}
```

### 15. What are Go packages and how do you import them?

**Answer:** Go's module system for organizing code. Import with `import` statement.
```go
import (
    "fmt"
    "net/http"
)
```

---

## 🔧 Intermediate Level (Questions 16-35)

### 16. What are goroutines and how do they work?

**Answer:** Lightweight threads managed by Go runtime. Much cheaper than OS threads.
```go
go function()  // start goroutine
go func() { fmt.Println("Hello") }()  // anonymous
```

### 17. How does the goroutine scheduler work?

**Answer:** M:N scheduler with work-stealing algorithm:
- **G**: Goroutines (user-space threads)
- **M**: OS threads  
- **P**: Logical processors
- Cooperative scheduling with preemption

### 18. What are channels and how do you use them?

**Answer:** Communication mechanism between goroutines. Send with `<-`, receive with `<-`.
```go
ch := make(chan int)
ch <- 42          // send
value := <-ch     // receive
```

### 19. What's the difference between buffered and unbuffered channels?

**Answer:**
- **Unbuffered**: Synchronous, sender blocks until receiver ready
- **Buffered**: Asynchronous up to buffer size, sender blocks when full

### 20. How do you use `select` statements?

**Answer:**
```go
select {
case msg := <-ch1:
    // handle msg
case <-timeout:
    // timeout
default:
    // non-blocking
}
```

### 21. What are interfaces in Go?

**Answer:** Define method signatures. Types implement interfaces implicitly.
```go
type Writer interface {
    Write([]byte) (int, error)
}
```

### 22. How does `interface{}` work and what are its performance impacts?

**Answer:** Empty interface holds any type. Performance impacts:
- Heap allocations (boxing)
- Type assertion overhead  
- Larger memory footprint (16 bytes)
- Dynamic dispatch

### 23. What are race conditions and how do you detect them?

**Answer:** Multiple goroutines accessing shared data with at least one write.
```bash
go run -race main.go  # Race detector
go test -race
```

### 24. How do you use sync.Mutex and sync.RWMutex?

**Answer:**
```go
var mu sync.Mutex
mu.Lock()
defer mu.Unlock()  // critical section

var rwMu sync.RWMutex
rwMu.RLock()    // read lock (multiple readers)
rwMu.Lock()     // write lock (exclusive)
```

### 25. What is sync.WaitGroup and when do you use it?

**Answer:** Waits for collection of goroutines to finish executing.
```go
var wg sync.WaitGroup
wg.Add(1)
go func() { defer wg.Done(); /* work */ }()
wg.Wait()
```

### 26. How does garbage collection work in Go?

**Answer:** Concurrent, tri-color mark-and-sweep collector. Runs when heap doubles (GOGC=100). Minimizes stop-the-world pauses.

### 27. What is escape analysis?

**Answer:** Compiler determines if variables go on stack or heap. Variables escape to heap when:
- Returned as pointer
- Stored in interface{}
- Captured by closures
- Too large for stack

### 28. How do you handle panics and recovery?

**Answer:**
```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
panic("error")
```

### 29. What are Go modules and how do you use them?

**Answer:**
```bash
go mod init example.com/project
go get github.com/package@v1.2.3
go mod tidy  # clean up
```

### 30. How do you write benchmarks in Go?

**Answer:**
```go
func BenchmarkFunction(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // code to benchmark
    }
}
// Run: go test -bench=.
```

### 31. What are build tags and how do you use them?

**Answer:**
```go
//go:build linux && !windows
package main
// Build: go build -tags "dev,debug"
```

### 32. How do you use context.Context?

**Answer:**
```go
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()

select {
case <-ctx.Done():
    return ctx.Err()
default:
    // work
}
```

### 33. What are Go's memory alignment rules?

**Answer:** Fields aligned to natural boundaries. Struct size is multiple of largest field alignment. Order fields largest to smallest for optimal layout.

### 34. How do you implement custom JSON marshaling?

**Answer:**
```go
func (p Person) MarshalJSON() ([]byte, error) {
    return json.Marshal(map[string]interface{}{
        "name": p.Name,
        "age":  p.Age,
    })
}

func (p *Person) UnmarshalJSON(data []byte) error {
    // custom unmarshaling logic
}
```

### 35. What is the difference between `new()` and `make()`?

**Answer:**
- `new(T)`: Allocates zeroed memory, returns `*T`
- `make(T)`: Creates and initializes slices/maps/channels, returns `T`

---

## 🚀 Advanced Level (Questions 36-50)

### 36. How does Go's memory model work?

**Answer:** Based on happens-before relationships. No guarantees across goroutines without synchronization. Use channels, mutexes, or atomic operations.

### 37. How do you implement an efficient goroutine pool?

**Answer:**
```go
type Pool struct {
    workers   int
    taskQueue chan func()
    wg        sync.WaitGroup
}
// Control worker count, handle backpressure, graceful shutdown
```

### 38. What are the best practices for using buffered channels?

**Answer:**
- Size based on system constraints
- Handle full buffers with select/default
- Avoid len()/cap() for flow control
- Monitor buffer utilization

### 39. How do you optimize Go programs for low latency?

**Answer:**
- Reduce allocations (object pooling)
- Tune GC (GOGC environment variable)
- Optimize memory layout
- Use atomic operations
- Profile hot paths

### 40. What are the different types of profiling in Go?

**Answer:**
- CPU: `go tool pprof`
- Memory: heap/alloc profiling
- Goroutine: deadlock detection
- Mutex: lock contention
- Block: synchronization blocking
- Trace: `go tool trace`

### 41. How do you implement custom atomic operations?

**Answer:**
```go
import "sync/atomic"

var counter int64
atomic.AddInt64(&counter, 1)
atomic.LoadInt64(&counter)
atomic.CompareAndSwapInt64(&counter, old, new)
```

### 42. What are the internals of Go's map implementation?

**Answer:** Hash table with buckets (8 key-value pairs each), overflow buckets for collisions, grows by doubling, not thread-safe.

### 43. How do you implement zero-copy optimizations?

**Answer:**
- Use `io.Copy()`
- Memory-mapped files
- `unsafe` package for type conversion
- Avoid interface{} boxing
- Efficient byte slice usage

### 44. What are Go's calling conventions and how do they affect performance?

**Answer:** Stack-based parameter passing, return values on stack, escape analysis determines allocation location.

### 45. How do you implement custom reflection in Go?

**Answer:**
```go
import "reflect"

t := reflect.TypeOf(obj)
v := reflect.ValueOf(obj)
// Dynamic type inspection and manipulation
```

### 46. What are the internals of Go's scheduler?

**Answer:** Work-stealing scheduler with global/local run queues, handles system calls by detaching M from P, preemptive scheduling via signals.

### 47. How do you implement lock-free data structures?

**Answer:** Use atomic operations, compare-and-swap loops, handle ABA problem, memory ordering considerations.

### 48. What are advanced context patterns?

**Answer:**
- Context chaining with timeouts
- Value propagation for request-scoped data
- Cancellation hierarchies
- Custom context implementations

### 49. How do you implement custom memory allocators?

**Answer:**
- Use `unsafe` package
- Memory pools with `sync.Pool`
- Off-heap storage
- Manual memory management for critical paths

### 50. What are Go's assembly optimizations and when would you use them?

**Answer:** Hand-written assembly for:
- Critical performance paths
- SIMD instructions
- Cryptographic operations
- Math-heavy computations
Use `go tool objdump` to analyze generated assembly.

---

### Key Commands
```bash
go run -race main.go      # Race detection
go test -bench=. -benchmem # Benchmarking
go tool pprof cpu.prof    # Profiling
go build -gcflags="-m"    # Escape analysis
```

---