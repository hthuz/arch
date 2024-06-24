package main

import (
	"fmt"
	"time"
)

type MyMsg struct {
	msgCh chan string
}

func main() {

	msgCh := make(chan string)
	numCh := make(chan int, 1)
	msg2Ch := make(chan MyMsg)

	go setMsg(msgCh)
	go setNum(numCh)
	go getMsg2(msg2Ch)

	// fmt.Println(<-msgCh)
	// fmt.Println(<-numCh)
	select {
	// case <-msgCh:
	// 	fmt.Println("msgCh set")
	case <-numCh:
		fmt.Println("numCh set")
	// If there's another goroutine waiting for msg2Ch (getMsg2), this case will be executed. and msg2Ch will be set to MyMsg{msgCh}
	case msg2Ch <- MyMsg{msgCh}:
		fmt.Println("msg2 channel executed")
	}

}

func setMsg(msgCh chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("setMsg executed")
	msgCh <- "hello"
}

func setNum(numCh chan int) {
	time.Sleep(4 * time.Second)
	fmt.Println("setNum executed")
	numCh <- 123
}

func getMsg2(msg2Ch chan MyMsg) {
	time.Sleep(3 * time.Second)
	fmt.Println(<-(<-msg2Ch).msgCh)

}
