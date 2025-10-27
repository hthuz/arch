// RWLock 是一个简单的读写锁实现
package main

import (
	"fmt"
	"sync"
	"time"
)

type RWLock struct {
	mu      sync.Mutex // 用于保护内部状态
	readers int        // 当前正在读取的协程数
	writer  bool       // 是否有写者正在写
	cond    *sync.Cond // 条件变量，用于唤醒等待的读/写者
}

func NewRWLock() *RWLock {
	lock := &RWLock{}
	lock.cond = sync.NewCond(&lock.mu)
	return lock
}

// 加读锁
func (rw *RWLock) RLock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	// 如果有写者在写，就等待
	for rw.writer {
		rw.cond.Wait()
	}
	rw.readers++
}

// 释放读锁
func (rw *RWLock) RUnlock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	rw.readers--
	// 如果没有读者了，唤醒可能等待的写者
	if rw.readers == 0 {
		rw.cond.Broadcast()
	}
}

// 加写锁
func (rw *RWLock) Lock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	// 等待直到没有读者且没有其他写者
	for rw.writer || rw.readers > 0 {
		rw.cond.Wait()
	}
	rw.writer = true
}

// 释放写锁
func (rw *RWLock) Unlock() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	rw.writer = false
	rw.cond.Broadcast() // 唤醒等待的读或写
}

func main() {
	rw := NewRWLock()
	wg := sync.WaitGroup{}

	// 启动多个读者
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			rw.RLock()
			fmt.Printf("Reader %d is reading\n", id)
			rw.RUnlock()
			wg.Done()
		}(i)
	}

	time.Sleep(200 * time.Millisecond)
	// 启动一个写者
	wg.Add(1)
	go func() {
		rw.Lock()
		fmt.Println("Writer is writing")
		rw.Unlock()
		wg.Done()
	}()

	wg.Wait()
}
