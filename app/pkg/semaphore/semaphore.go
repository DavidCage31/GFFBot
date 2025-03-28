package semaphore

type Semaphore struct {
    c chan struct{}
}

func (s *Semaphore) Acquire() {
	s.c <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.c
}

func New(tickets int) *Semaphore {
	return &Semaphore{
		c: make(chan struct{}, tickets),
    }
}