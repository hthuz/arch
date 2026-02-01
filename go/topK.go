package main

import "fmt"

func main() {
	// 347
}

func topKFrequent(nums []int, k int) []int {

	count := make(map[int]int)
	for _, num := range nums {
		count[num] += 1
	}
	m := make(map[int][]int)
	counts := make([]int, 0)
	for k, v := range count {
		m[v] = append(m[v], k)
		counts = append(counts, v)
	}

	// heapify counts
	sift_down := func(idx int) {
		for idx != 0 {
			parent := (idx - 1) / 2
			if counts[idx] > counts[parent] {
				counts[idx], counts[parent] = counts[parent], counts[idx]
			} else {
				break
			}
			idx = parent
		}
	}
	for i := len(counts)/2 - 1; i >= 0; i-- {
		sift_down(i)
	}

	fmt.Println(counts)

	res := make([]int, 0)
	for _, count := range counts {
		if len(res) == k {
			break
		}
		if len(res)+len(m[count]) > k {
			res = append(res, m[count][:k-len(res)+1]...)
			break
		} else {
			res = append(res, m[count]...)
		}
	}

	return res
}
