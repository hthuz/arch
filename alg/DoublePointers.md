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



### binary search



### n sum problem



3 sum problem. Given `nums`, find all`nums[i], nums[j], nums[k]`(no repetitive) such that their sum is `0`

given `[-1, 0, 1, 2, -1, -4]` return

`[-1, -1, 2]`, `[-1, 0, 1]`

```go
func threeSum(nums []int) [][]int {
    slices.Sort(nums)
    res := make([][]int, 0)

    for i, _ := range nums {
        target := -nums[i]
        // 跳过i的重复，比如上述例子中两个-1
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
                
                // 跳过left的重复
                for left < right && nums[left-1] == nums[left] {
                    left++
                }
                // 跳过right的重复
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



找到满足xx条件的最xxx的区间/子串/子数组



- 设置窗口 [left, right) (一开始为[0, 0), 窗口为空)
- **寻找可行解**：右移right，直到窗口满足xx条件
- **寻找最优解**： 右移left， 直到窗口不满足xx条件，
- 直到right到达end



```go

left := 0
right := 0
for right < len(nums) {
    // move right, update window
    
    for (window satisfying condition) {
        // update global data
        
       	// move left, update window
    }
}

```



### 76. 最小覆盖子串

给你一个字符串 `s` 、一个字符串 `t` 。返回 `s` 中涵盖 `t` 所有字符的最小子串。如果 `s` 中不存在涵盖 `t` 所有字符的子串，则返回空字符串 `""` 。

```
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
```



```go
func minWindow(s string, t string) string {
    left := 0
    right := 0
    res := ""
    cur_str := ""
    for right < len(s) {
        cur_str = append(cur_str, s[right])
        for 
        
    }
    
}


```





### 3. 无重复字符的最长子串



给定一个字符串 `s` ，请你找出其中不含有重复字符的 **最长**  的长度。

```
输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。注意 "bca" 和 "cab" 也是正确答案。
```



```go
func lengthOfLongestSubstring(s string) int {
    ans := 0
    m := make(map[byte]struct{})
    left := 0
    right := 0

    exist := func(b byte) bool {
        if _, exist := m[b]; exist {
            return true
        }
        return false
    }
    for right < len(s) {
        for exist(s[right]) {
            delete(m, s[left])
            left++
        }
        m[s[right]] = struct{}{}
        ans = max(ans, len(m))
        right++
    }
    return ans

}
func lengthOfLongestSubstring(s string) int {
    window := [128]bool{}
    left := 0
    right := 0
    ans := 0
    for right < len(s) {
        for window[s[right]] {
            window[s[left]] = false
            left++
        }
        window[s[right]] = true
        right++
        ans = max(ans, right-left)
    }
    return ans
}

func lengthOfLongestSubstring(s string) int {
    window := [128]bool{}
    left := 0
    right := 0
    ans := 0
    for right, _ := range < s {
        for window[s[right]] {
            window[s[left]] = false
            left++
        }
        window[s[right]] = true
        ans = max(ans, right-left)
    }
    return ans
}
```



### 30. 串联所有单词的子串



### 159. 至多包含两个不同字符的最长子串



### 209. 长度最小的子数组



### 239. 滑动窗口最大值



### 567. 字符串的排列



### 632. 最小区间



### 727. 最小窗口子序列

