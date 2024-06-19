
package main

import (
	"fmt"
	"math"
	"errors"
)

func sqrt(num int) (float64, error) {
	if num < 0 {
		return float64(num), errors.New("num is less than 0")
	}
	return math.Sqrt(float64(num)), nil

}

func main() {
	var num int
	var sqrt_num float64
	num = -2
	sqrt_num, err := sqrt(num)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(sqrt_num)
}
