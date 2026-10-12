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
    // 
    // right更新的地方不一定在这，就我自己遇到的题目而言
    // 都是现在 for window satisfy condition 前，把right放入window, 然后for window satisfy condition 后
    // 更新right
    
    for (window satisfying condition) {
        // update global data
        
       	// move left, update window
    }
}

```



更准确的

如果是求满足xxx条件的最小xxx

```go
left := 0
right := 0

for right < len(nums) {
    window[right]++
    
    for (window satisfy condition) {
        ans = min(ans, right-left+1)
        
        
        window[left]--
        left += 1
    }
    
    right += 1
}
return ans
```



如果是求满足xx条件的最长xxx

这个遇到再总结吧，note一下先





### 76. 最小覆盖子串

给你一个字符串 `s` 、一个字符串 `t` 。返回 `s` 中涵盖 `t` 所有字符的最小子串。如果 `s` 中不存在涵盖 `t` 所有字符的子串，则返回空字符串 `""` 。

```
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
```



```go
func minWindow(s string, t string) string {
    

    t_count := make(map[byte]int)
    for i := range t {
        t_count[t[i]] += 1
    }

    min_len := math.MaxInt32

    left := 0
    right := 0

    s_count := make(map[byte]int)

    ans := ""


    satisfy := func() bool {

        for c, count := range t_count {
            if s_count[c] < count {
                return false
            }
        }
        return true

    }
    for right < len(s) {

        s_count[s[right]] += 1

        // shorten left/ inc left

        for satisfy() {
            if right-left+1 < min_len {
                min_len = right-left+1
                ans = s[left:right+1]
            }
			
            // 这里更新过后，可能就已经不satisfy窗口的要求了，但是已经把ans记录了，最后跳出这个for就是不满足的一个情况
            s_count[s[left]] -= 1
            left += 1
        }
        
        right++
    }

    return ans
    
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

给定一个含有 `n` 个正整数的数组和一个正整数 `target` **。**找出该数组中满足其总和大于等于 `target` 的长度最小的 **子数组**

```go
func minSubArrayLen(target int, nums []int) int {

	left, right := 0, 0
	sum := 0
	ans := math.MaxInt32

	for right < len(nums) {
		sum += nums[right]

		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left += 1
		}

		right += 1

	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans

}
```





### 239. 滑动窗口最大值



### 567. 字符串的排列

给你两个字符串 `s1` 和 `s2` ，写一个函数来判断 `s2` 是否包含 `s1` 的 。如果是，返回 `true` ；否则，返回 `false` 。

换句话说，`s1` 的排列之一是 `s2` 的 **子串** 。

```
输入：s1 = "ab" s2 = "eidbaooo"
输出：true
解释：s2 包含 s1 的排列之一 ("ba").
```

```go
// 自己题解，非官方题解
func checkInclusion(s1 string, s2 string) bool {

    s1_count := [26]int{}
    s2_count := [26]int{}
    for _, s := range s1 {
        s1_count[s-'a'] += 1
    }
    l := 0
    r := 0
    for r < len(s2) {
        s2_count[s2[r]-'a'] += 1
        for s2_count[s2[r]-'a'] > s1_count[s2[r]-'a'] {
            s2_count[s2[l]-'a'] -= 1
            l += 1
        }
        if s2_count[s2[r]-'a'] == s1_count[s2[r]-'a'] {
            // check
            equal := true
            for i := range 26 {
                if s1_count[i] != s2_count[i] {
                    equal = false
                    break
                }
            }
            if equal {
                return true
            }
        }
        r += 1
        
    }
    return false

}
```



### 632. 最小区间



### 727. 最小窗口子序列



### 3095 或值至少为 K 的最短子数组

```go
func minimumSubarrayLength(nums []int, k int) int {

    if k == 0 {
        return 1
    }
    ans := len(nums) + 1

    left := 0
    right := 0
    window := make([]int, 0)
    for right < len(nums) {
        window = append(window, nums[right])

        for getVal(window) >= k {
            ans = min(ans, right - left + 1)
            
            window = window[1:]
            left++
        }

        right++
    } 
    if ans == len(nums) + 1 {
        return -1
    }
    return ans
}

func getVal(nums []int) int {
    ans := 0
    for _, x := range nums {
        ans = ans | x
    }
    return ans
}
```

