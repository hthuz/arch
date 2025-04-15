// Golang program to illustrate
// reflect.CanInterface() Function

package main

import (
	"fmt"
	"reflect"
)

// Main function
func main() {

	string := "ABC"
	meta := reflect.ValueOf(&string)

	// use of CanInterface() method
	fmt.Println("canInterface:", meta.CanInterface())
}
