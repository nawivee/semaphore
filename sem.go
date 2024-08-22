package sem

import "sync"

type Sem struct {
	wg  sync.WaitGroup
	sem chan struct{}
}

func New(limit uint) *Sem {
	return &Sem{sem: make(chan struct{}, limit)}
}

func (s *Sem) Acquire() {
	s.wg.Add(1)
	s.sem <- struct{}{}
}

func (s *Sem) Wait() {
	s.wg.Wait()
}
func (s *Sem) Release() {
	<-s.sem
	s.wg.Done()
}
