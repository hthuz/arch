package main

import (
	"fmt"
	"math"
)

func main() {

	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	// 一开始是 左右分的情况是
	// 1 | 2
	// 3 | 4
	// 后面分的情况是
	// 1 2 |
	// | 3 4
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println("res", res)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// nums1 按照 i切分， nums2 按照 j 切分
	// nums1[i-1] <= nums2[j], nums2[j-1] <= nums1[i]
	// if odd, max of left part
	// if even, median of max of left part and min of right part
	if len(nums2) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	left, right := 0, m

	for left < right {
		i := (left + right + 1) / 2
		j := (m+n+1)/2 - i

		fmt.Println("i", i, "j", j)

		if nums1[i-1] > nums2[j] {
			// i过大了， i需要往左移
			right = i - 1
		} else {
			left = i
		}

	}
	i, j := left, (m+n+1)/2-left

	var nums1Left, nums1Right int
	var nums2Left, nums2Right int

	// 这些都是边界情况处理
	if i > 0 {
		nums1Left = nums1[i-1]
	} else {
		nums1Left = math.MinInt32
	}

	if i < m {
		nums1Right = nums1[i]
	} else {
		nums1Right = math.MaxInt32
	}

	if j > 0 {
		nums2Left = nums2[j-1]
	} else {
		nums2Left = math.MinInt32
	}

	if j < n {
		nums2Right = nums2[j]
	} else {
		nums2Right = math.MaxInt32
	}

	maxLeft := max(nums1Left, nums2Left)
	minRight := min(nums1Right, nums2Right)

	// maxLeft := max(nums1[i-1], nums2[j-1])
	// minRight := min(nums1[i], nums2[j])

	if (m+n)%2 == 1 {
		return float64(maxLeft)
	}

	return float64(maxLeft+minRight) / 2

}
