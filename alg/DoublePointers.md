# Double Pointers



## Fast or slow pointers

Examples 

- cycles in linked list
- duplicate elements in array/list

### Example of removing duplicates from a sorted array

Given `nums []int`, remove duplicates in place from it and return the length of updated array
For example, Given [1,1,2], return [1,2, _], with length 2

A naive approach is to evaluate 
`nums[i] == nums[i+1]`, 
if true, remove `nums[i+1]`, keep `i` the same
else, update `i`

The idea is to use `slow` ptr keeping track of return array  
use `fast` ptr finding deduplicated elements

A problem is whether use `nums[fast] == nums[fast+1]` or `nums[fast] == nums[fast-1]` as duplication condition


```go
func removeDuplicates(nums []int) int {

    if len(nums) == 0 {
        return 0
    }
    // keep the first one
    slow := 1
    for fast := 0; fast < len(nums) - 1; fast++ {
        // meet a new element, put new elemnt
        if nums[fast] != nums[fast+1] {
            nums[slow] = nums[fast+1]
            slow++
        }
    }
    return slow
}

```



## Left Right Pointers

Left pointers move to the right and right pointers move to the left

Usually on ordered array/list

Examples

- binary search
- n sum problem



3 sum problem. Given `nums`, find all`nums[i], nums[j], nums[k]`(no repetitive) such that their sum is `0`

given `[-1, 0, 1, 2, -1, -4]` return

`[-1, -1, 2]`, `[-1, 0, 1]`

```go
func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    res := make([][]int, 0)

    for i, _ := range nums {
        target := -nums[i]
        if i - 1 >= 0 && nums[i-1] == nums[i] {
            continue
        }
        left := i + 1
        right := len(nums) - 1
        for left < right {
            sum := nums[left] + nums[right]
            if sum < target {
                left++
            } else if sum > target {
                right--
            } else {
                res = append(res, []int{nums[i], nums[left], nums[right]})
                left++
                right--
                for left < right && nums[left-1] == nums[left] {
                    left++
                }
                for left < right && nums[right] == nums[right + 1] {
                    right--
                }
                
            }
        }
    }
    return res

}
```






## Sliding Window



