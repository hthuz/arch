# permutation

## 全排列

给定一个数组，每个元素都不相同，返回数组全排列 => 回溯

给定一个数组，元素可能相同，返回数组的全排列， （[1,1,2 ] 和 [1,1,2] 算同一种）=> 回溯，同时剪枝



## 下一个排序

根据字典序，比如给[1,2,3]， 下一个是 [1,3,2], 给 [2,1,3]， 下一个是[2,3,1]   => 遍历，没什么技巧，主要就是记住吧



比如给的是 1,6,5,2,4,3 前面有一部分肯定不用变

如果nums没有重复，可以下面这样子写，但是如果nums有重复，需要更复杂

```go
func nextPermutation(nums []int)  {
   
    // find smaller
    pivot := len(nums)-2
    for pivot >= 0 {
        if nums[pivot] < nums[pivot + 1] {
            break
        }
        pivot--
    }
    
    // 如果smaller没有找到，直接reverse, 要么全数组都是降序，要么len(nums) < 2
    if pivot < 0 {
        reverse(nums)
        return
	}
    
    
    // find larger
    swapped := len(nums) - 1
    // 如果没有重复，可以直接
    // for nums[swapped] < nums[pivot] {
    //     swapped -= 1
    // } 
    // 如果重复，比如[1,5,1] -> [5,1,1]， pivot在0, 应该让swapped定位在5,而不是1
    for nums[swapped] <= nums[pivot] {
        swapped -= 1
    }
    nums[pivot], nums[swapped] = nums[swapped], nums[pivot]
    
    // 现在nums[pivot+1:]一定是降序的
    // reverse
    reverse(nums[pivot+1:])

}

func reverse(nums []int) {
    for i := 0; i < len(nums) / 2; i++ {
        j := len(nums)-1-i
        nums[i], nums[j] = nums[j], nums[i]
    }
}

```



