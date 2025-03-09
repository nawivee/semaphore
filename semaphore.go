// Package semaphore provides a simple implementation of a counting semaphore
// that limits concurrent access to resources.
package semaphore

import "sync"

// Semaphore implements a counting semaphore pattern to limit concurrent operations.
// It combines the functionality of a channel-based semaphore with a wait group
// to provide synchronization capabilities.
type Semaphore struct {
	waitGroup sync.WaitGroup
	tokens    chan struct{}
}

// New creates a new semaphore with the specified concurrency limit.
// The limit defines the maximum number of concurrent operations allowed.
func New(concurrencyLimit uint) *Semaphore {
	return &Semaphore{
		tokens: make(chan struct{}, concurrencyLimit),
	}
}

// Acquire acquires a token from the semaphore.
// If no tokens are available, the call will block until one becomes available.
func (s *Semaphore) Acquire() {
	s.waitGroup.Add(1)
	s.tokens <- struct{}{}
}

// Release releases a token back to the semaphore and signals completion of a task.
// This method should be called exactly once for each Acquire call.
func (s *Semaphore) Release() {
	<-s.tokens
	s.waitGroup.Done()
}

// Wait blocks until all acquired tokens have been released.
// This allows for synchronization of all concurrent operations.
func (s *Semaphore) Wait() {
	s.waitGroup.Wait()
}
