package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

type Winsize struct {
	ws_row    uint16
	ws_col    uint16
	ws_xpixel uint16
	ws_ypixel uint16
}

func main() {

	// syscall.
	var win Winsize
	win_ptr := uintptr(unsafe.Pointer(&win))

	// syscall.Syscall(uintptr(syscall.Stdin))
	syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), syscall.TIOCGWINSZ, win_ptr)

	fmt.Println(win)

}
