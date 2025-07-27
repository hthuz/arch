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




## Sliding Window



