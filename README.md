# Semaphore

A lightweight Go package that provides a simple, efficient implementation of a counting semaphore with synchronization capabilities.

## Overview

The `semaphore` package offers a straightforward way to limit concurrent operations in Go applications. It combines a channel-based semaphore with a wait group to both limit concurrency and provide synchronization of concurrent tasks.

## Installation

```bash
go get github.com/nawivee/semaphore
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "time"
    
    "github.com/nawivee/semaphore"
)

func main() {
    // Create a semaphore with a concurrency limit of 3
    sem := semaphore.NewSemaphore(3)
    
    // Launch 10 goroutines
    for i := 0; i < 10; i++ {
        // Acquire a token from the semaphore
        sem.Acquire()
        go func(id int) {    
            // Ensure token is released when done
            defer sem.Release()
            
            // Simulate work
            fmt.Printf("Worker %d starting\n", id)
            time.Sleep(2 * time.Second)
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }
    
    // Wait for all goroutines to complete
    sem.Wait()
    fmt.Println("All workers have completed")
}
```

### Use Case: Rate Limiting HTTP Requests

```go
func fetchURLs(urls []string, concurrencyLimit uint) []string {
    sem := semaphore.NewSemaphore(concurrencyLimit)
    results := make([]string, len(urls))
    
    for i, url := range urls {
        sem.Acquire()
        go func(index int, url string) {            
            defer sem.Release()
            
            // Fetch the URL
            resp, err := http.Get(url)
            if err != nil {
                results[index] = fmt.Sprintf("Error: %v", err)
                return
            }
            defer resp.Body.Close()
            
            // Process the response
            body, _ := io.ReadAll(resp.Body)
            results[index] = string(body)
        }(i, url)
    }
    
    // Wait for all requests to complete
    sem.Wait()
    return results
}
```

## Testing

Run the tests using:

```bash
go test -v ./...
```

## Best Practices

- Always pair `Acquire()` with `Release()` calls, preferably using `defer` to ensure tokens are released even if a panic occurs.
- Use a concurrency limit appropriate for your resources. For CPU-bound tasks, consider using `runtime.NumCPU()` as a guideline.
- For I/O-bound tasks like network requests, higher concurrency limits may be appropriate.

## Thread Safety

All methods in this package are thread-safe and can be safely called from multiple goroutines.

## License

MIT