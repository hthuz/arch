package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}

func findDuplicate(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum - len(nums)*(len(nums)-1)/2
}
