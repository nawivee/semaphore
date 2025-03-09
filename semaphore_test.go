package semaphore

import (
	"fmt"
	"testing"
	"time"
)

func TestSem(t *testing.T) {

	count := 0
	sem := New(2)
	go func() {
		sem.Acquire()
		count++
		sem.Acquire()
		count++
		sem.Acquire() // this and next acquires should be blocked
		count++
		sem.Acquire()
		count++
	}()

	time.Sleep(100 * time.Microsecond)
	if count != 2 {
		t.Errorf("count should be 2 but got %d", count)
	}
	sem.Release()
	time.Sleep(100 * time.Microsecond)
	if count != 3 {
		t.Errorf("count should be 3 but got %d", count)
	}

}

func ExampleSemaphore() {
	// Create a semaphore with concurrency limit of 2
	sem := New(2)

	// Launch 4 goroutines that will try to acquire the semaphore
	for i := 0; i < 4; i++ {
		sem.Acquire()
		go func(id int) {
			defer sem.Release()

			// Simulate some work
			fmt.Printf("Worker %d is working\n", id)
			time.Sleep(100 * time.Millisecond)
		}(i)

		// to ensure order of execution
		// you don't need this in your code
		// it's required for example
		// to have predictable output
		time.Sleep(100 * time.Millisecond)
	}

	// Wait for all goroutines to complete
	sem.Wait()

	// Output:
	// Worker 0 is working
	// Worker 1 is working
	// Worker 2 is working
	// Worker 3 is working
}

// TODO: Add Readme
