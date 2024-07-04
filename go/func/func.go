
package main

import (
	"fmt"
	"math"
	"errors"
)

func sqrt(num int) (float64, error) {
	if num < 0 {
		err := fmt.Sprintf("%d is less than 0", num)
		return float64(num), errors.New(err)
	}
	return math.Sqrt(float64(num)), nil

}

func main() {

	// Test
	var num int
	var sqrt_num float64
	num = -2
	sqrt_num, err := sqrt(num)
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println(sqrt_num)
	}

	// For loop
	for i := -5; i < 5; i++ {
		sqrt_num, err := sqrt(i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(sqrt_num)
	}

	// Slice
	nums := [] int {-1,-2,-3,1,2,3}
	fmt.Println(nums)
	for k,v := range nums {
		fmt.Println(k,v)
	}



}
