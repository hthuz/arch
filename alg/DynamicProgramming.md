
# Dynamic Programming


## Minimum path sum

The most basic template for traversing a 2D dp array
Leetcode 64, given a grid, find a way with minimum sum of values, by only going down or going right

e.g.
If grid is
1 3 1
1 5 1
4 2 1

Then dp array is

1 4 5
2 7 6
6 8 7

dp(i,j) = minimum sum of values by going to position (i,j)

dp(i,j) = min(dp(i-1,j), dp(i, j-1)) + val(i,j)

```go
func minPathSum(grid [][]int) int {
    height := len(grid)
    width := len(grid[0])
    dp := make([][]int, height)
    for i := range height {
        dp[i] = make([]int, width)
    }
    // initialize
    dp[0][0] = grid[0][0]
    for i := 1; i < width; i++{
        dp[0][i] = dp[0][i-1] + grid[0][i]
    }
    for i := 1; i < height; i++ {
        dp[i][0] = dp[i-1][0] + grid[i][0]
    }
    
    // dp
    for i := 1; i < height; i++ {
        for j := 1; j < width; j++ {
            dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
        }
    }
    return dp[height-1][width-1]
    
}
```



## MaxSubArray

max sum of a continuous array
The naive approch is list all subarrays and compute the largest one.
It's equilavent to list a dp(start, end), which records sum of array(start, end), and then find the best one

```go
func maxSubArray(nums []int) int {
    res := nums[0]
    for i := range len(nums) {
        for j := i; j < len(nums); j++ {
            s := 0
            for k := i; k <= j; k++ {
                s += nums[k]
            }
            res = max(s, res)
        }
    }
    return res
}
```

A more wise way is to define dp(n) as maxSubArraySum with n as tail of subarray

then, dp(n) =
dp(n-1) + nums(n) if include subarraySum from previous
or nums(n) if start a new head,
depending on which one is larger


And result is max(dp)

```go
func maxSubArray(nums []int) int {
    n := len(nums)
    m := nums[0]
    dp := make([]int, n)
    dp[0] = nums[0]
    for i := 1; i < n; i++ {
        dp[i] = max(nums[i] + dp[i-1], nums[i]) 
        m = max(m, dp[i])
    }
    return m
}

// remove array
func maxSubArray(nums []int) int {
    prev := nums[0]
    res := nums[0]
    for i := 1; i < len(nums); i++ {
        cur := max(nums[i] + prev, nums[i])
        res = max(res, cur)
        prev = cur
    }
    return res
}

// remove prev and cur
func maxSubArray(nums []int) int {
    sum := nums[0]
    res := nums[0]
    for i := 1; i < len(nums); i++ {
        sum = max(nums[i] + sum, nums[i])
        res = max(res, sum)
    }
    return res
}

```

## longestPalindrome

dp(i,j) is denoted as the length of longest palindrome substring of string s(i..j) 
(for all substrings of s(i,j), its longest palindrome length is dp(i,j))

Then
if s(i) == s(j) && s(i+1,j-1) is palindrome
dp(i,j) = dp(i+1, j-1) + 2
else 
dp(i,j) = max(dp(i+1, j-1), dp(i, j-1), dp(i+1, j))

Initial condition: 
all dp(i,i) = 1 
all dp(i, i+1) = 2 if s(i) == s(i+1) else 1




## maxProfit

leetcode 121
Given prices, buying at day i and selling at day j, find out maxprofit prices(j) - prices(i)
define dp(i) as max profit of selling at day i, answer is max(dp)
dp(i) = dp(i-1) - prices(i-1) + prices(i) if keep the same buying date
else dp(i) = 0 if choose new buying date

```go

func maxProfit(prices []int) int {
	prev := 0
    res := 0
	for i := 1; i < len(prices); i++ {
        cur := max(0, prev-prices[i-1] + prices[i])
        res = max(res, cur)
        prev = cur
	}

	return res
}
```


In a more abstract way
```go

func maxProfit(prices []int) int {
    profit := 0
    res := 0
	for i := 1; i < len(prices); i++ {
        profit = max(0, profit - prices[i-1] + prices[i])
        res = max(res, profit)
	}
	return res
}
```

prices := (8,5,6,4,9)
delta = (-3,1,-2,5)

This is actually equilvaent/similar to computing maxsubArray sum of delta, 

if dp(i) = max(dp(i-1), dp(i-1) - prices(i-1) + prices(i)), it becomes leetcode 122. 
Max profit by holding at most one stock
dp(i) = max(dp(i-1), dp(i-1) + delta(i))





## Jumping game

给定一个长度为 `n` 的 **0 索引**整数数组 `nums`。初始位置在下标 0。

每个元素 `nums[i]` 表示从索引 `i` 向后跳转的最大长度。换句话说，如果你在索引 `i` 处，你可以跳转到任意 `(i + j)` 处：

- `0 <= j <= nums[i]` 且
- `i + j < n`

返回到达 `n - 1` 的最小跳跃次数。测试用例保证可以到达 `n - 1`。



e.g. given `nums = [2,3,1,1,4]`, return 2

```go
func jump(nums []int) int {
    dp := make([]int, len(nums))
 	
    for i := 1; i < len(nums); i++ {
        min_val := len(nums)
        for j := 0; j < i; j++ {
            if nums[j] + j >= i {
                min_val = min(min_val, dp[j] + 1)
            }
        }
        dp[i] = min_val
    }
    return dp[len(nums)-1]
}
```

以上方法复杂度仍会达到`O(n^2)`



A better method

记录能达到的最远边界, 贪心

```go
func jump(nums []int) int {
    farthest := 0
    cur_position := 0
    steps := 0
    // 最后一个位置不需要
    for i := 0; i < len(nums)-1; i++ {
        farthest = max(farthest, nums[i] + i)
        // 只有在你到达你当前能跳到最远的地方时，你才跳
        if cur_position == i {
            steps++
            // 你跳到了你现在能到达的最远的地方
            cur_position = farthest
        }
    }
    return steps
}w
```



