package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// IncreRaw may have repetition or skip some numbers
// e.g. 160 161 161 163 164 165 166
// or 30 31 32 36 36 36 36 37 38
// IncreAtomic has every occurence of number but the order may not be correct
// e.g. 61 62 63 64 65 68 66 67 69 70 71 72 73
// IncreLock and IntreLockAndAtomic always give the continuous result

var counter int64
var list []int64
var mu sync.Mutex
var IncreTimes = 50
var WorkerNum = 4
var TestNum = 10

type Tester struct {
	WorkerFunc func(*sync.WaitGroup)
}

func (t *Tester) testOneTime() (bool, int64) {
	counter = 0
	list = []int64{}
	var wg sync.WaitGroup
	wg.Add(WorkerNum)
	for range WorkerNum {
		go t.WorkerFunc(&wg)
	}
	wg.Wait()
	if ok, pos := isContinuous(list); !ok {
		fmt.Println(list)
		fmt.Println(pos)
	}
	return isContinuous(list)
}

func (t *Tester) Test() bool {
	fmt.Println(runtime.FuncForPC(reflect.ValueOf(t.WorkerFunc).Pointer()).Name())
	testRes := true
	for range TestNum {
		res, _ := t.testOneTime()
		testRes = testRes && res
	}
	return testRes
}

func main() {
	tester := Tester{}
	// tester.WorkerFunc = increRaw
	// fmt.Println(tester.Test())
	// tester.WorkerFunc = increLock
	// fmt.Println(tester.Test())
	tester.WorkerFunc = increAtomic
	fmt.Println(tester.Test())
	// tester.WorkerFunc = increLockAndAtomic
	// fmt.Println(tester.Test())
}

func increRaw(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < IncreTimes; i++ {
		counter += 1
		mu.Lock()
		list = append(list, counter)
		mu.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
}

func increLock(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < IncreTimes; i++ {
		mu.Lock()
		counter += 1
		list = append(list, counter)
		mu.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
}

func increAtomic(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < IncreTimes; i++ {
		atomic.AddInt64(&counter, 1)
		mu.Lock()
		list = append(list, counter)
		mu.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
}

func increLockAndAtomic(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < IncreTimes; i++ {
		mu.Lock()
		atomic.AddInt64(&counter, 1)
		list = append(list, counter)
		mu.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
}

// If not continuous, return false and the position that's not continuous
func isContinuous(arr []int64) (bool, int64) {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] != arr[i]+1 {
			return false, arr[i]
		}
	}
	return true, 0
}
