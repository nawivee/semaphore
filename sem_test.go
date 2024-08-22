package sem

import (
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

}

// TODO: Add example how to use
// TODO: Add Readme
