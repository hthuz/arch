package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	url          = "https://bsc-dataseed2.ninicoin.io"
	method       = "eth_getBlockByNumber"
	startBlock   = 0
	endBlock     = 1000
	TimeInterval = 3
	batchSize    = 5
)

func main() {
	// SimSyncRequests()
	SimAsyncRequests()
}

func SimSyncRequests() {
	ticker := time.NewTicker(TimeInterval * time.Second)
	reqCount := 0
	go func() {
		for range ticker.C {
			fmt.Printf("req %d for %d seconds\n", reqCount, TimeInterval)
			reqCount = 0
		}
	}()

	for block := startBlock; block < endBlock; block++ {
		params := []interface{}{IntToHex(block), false}
		resp, err := SendJSONRPCRequest(url, method, params)
		reqCount++
		if err != nil {
			fmt.Println("Error:", err)
			fmt.Println(resp.Result)
			continue
		}
		// fmt.Println(resp.Result)
	}
}
func SimAsyncRequests() {

	ticker := time.NewTicker(TimeInterval * time.Second)
	reqCount := 0
	var reqCountLock sync.Mutex
	go func() {
		for range ticker.C {
			fmt.Printf("req %d for %d seconds\n", reqCount, TimeInterval)
			reqCountLock.Lock()
			reqCount = 0
			reqCountLock.Unlock()
		}
	}()

	resps := make(chan BlockResponse, batchSize)
	var wg sync.WaitGroup
	for base := startBlock; base < endBlock; base += batchSize {
		fmt.Println("loop")
		// Send batchSize requests and process them asynchronously
		for block := base; block < base+batchSize; block++ {
			wg.Add(1)
			go func(wg *sync.WaitGroup, block int) {
				defer wg.Done()
				params := []interface{}{IntToHex(block), false}
				resp, err := SendJSONRPCRequest(url, method, params)
				reqCountLock.Lock()
				reqCount++
				reqCountLock.Unlock()
				if err != nil {
					fmt.Println("Error:", err)
				}
				resps <- resp.Result
			}(&wg, block)

		}
		go func() {
			wg.Wait()
			close(resps)
		}()

		for resp := range resps {
			fmt.Println(resp)
		}
	}

}

func IntToHex(i int) string {
	return fmt.Sprintf("0x%x", i)
}
