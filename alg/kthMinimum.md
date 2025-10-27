# kthMinimum



## K th minimum of an array

```go
// k starts with 0
func quick_select(nums []int, k int) int {
	pivot := nums[0]
	smaller := []int{}
	bigger := []int{}
	pivot_num := 0
	for _, num := range nums {
		if num < pivot {
			smaller = append(smaller, num)
			continue
		}
		if num > pivot {
			bigger = append(bigger, num)
			continue
		}
		pivot_num += 1
	}

	// NOTE here: 
	// using k >= len(smaller) or k > len(smaller)
	// using k < len(smaller)+pivot or k < len(smaller) + pivot_num 
	if k >= len(smaller) && k < len(smaller)+pivot_num {
		return pivot
	}
	if k < len(smaller) {
		return quick_select(smaller, k)
	}
	return quick_select(bigger, k-len(smaller)-pivot_num)
}
```



## kth minimum of two sorted array

