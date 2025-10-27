package lock

import "sync"

type RWLock struct {
	mu      sync.Mutex
	readers int
	writer  bool
	cond    *sync.Cond
}

func (rw *RWLock) RLock() {

	rw.mu.Lock()
	defer rw.mu.Unlock()

	for rw.writer {
		rw.cond.Wait()
	}
	rw.readers++

}

func (rw *RWLock) RUnlock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	rw.readers--

	if rw.readers == 0 {
		rw.cond.Broadcast()
	}

}

func (rw *RWLock) Lock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	for rw.writer || rw.readers > 0 {
		rw.cond.Wait()
	}
	rw.writer = true
}

func (rw *RWLock) Unlock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	rw.writer = false
	rw.cond.Broadcast()

}
