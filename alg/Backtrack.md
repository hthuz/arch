# Backtrack



两个常见的情况：

如果元素有重复，怎么剪枝

或者原数组元素不重复，但是元素可以选多次，也会有重复。





https://zhuanlan.zhihu.com/p/93530380

Examples include full permutation and N-queen

For each node, it has a **path** and **choices**

and you need to note about the **end condition**

![backtrack](/home/autentico/arch/alg/img/backtrack.png)



```python
for choice in choices:
	# make a choice
	path.add(choice)
	backtrack(path, choices)
	# undo choice
	path.remove(choice)
```



```python
# Template
result = []
def backtrack(path, choices):
    if satisfy_end_condition:
        result.add(path)
        return

    for choice in choices:
        # make a valid choice
        backtrack(path, choices)
        # undo choice
```



## 全排列

```go
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	var dfs func(int)

	dfs = func(i int) {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := range len(nums) {
			if used[i] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true

			dfs(i)

			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(0)
	return ans

}

```



Given N numbers

Number of nodes in the first layer 1

Number of nodes in the second layer N

Number of nodes in the third layer N(N-1)

Number of leaves:  N!

Time complexity: N!



## 组合总和

Given `candidates`, each element unique, find all combinations that sum to `target` (leetcode 39)
you can use the same element for multiple times.

For example, candidates: [2,3,6,7], target: 7
The problem is that there may be repetition, for example, [2,2,3], [2,3,2], [3,2,2] so you need to check for repetition, which is inefficient

A wiser way


```go
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)

	var dfs func(int, int)

	dfs = func(i int, sum int) {
		if i == len(candidates) {
			return
		}
		if sum > target {
			return
		}
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}

		dfs(i+1, sum)

		path = append(path, candidates[i])
		dfs(i, sum+candidates[i])
		path = path[:len(path)-1]

	}
	dfs(0, 0)
	return ans
}

```





## trick

在回溯时，上传结果是有时候需要创建一个copy

```go
temp := make([]int, len(path))
copy(temp, path)
res = append(res, temp)
```

这种方式的另一个简单写法是

```go
res = append(res, append([]int(nil), path...)) 
// or 
res = append(re(res, append([]int(nil), path...)) 
// or 
res = append(ress, append([]int{}, path...))
```





## Nqueen

```go
func solveNQueens(n int) [][]string {

	res := make([][]string, 0)
	path := make([]string, 0)
	backtrack(n, path, &res)
	return res
}

// path is current board
// choices is all columns in new row
func backtrack(n int, path []string, res *[][]string) {
	if len(path) == n {
		temp := make([]string, n)
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for choice := range n {
		newRow := make([]byte, n)
		for j := range n {
			newRow[j] = '.'
		}
		newRow[choice] = 'Q'
		path = append(path, string(newRow))
        
		if isvalid(path) {
			backtrack(n, path, res)
		}
		path = path[:len(path)-1]
	}
}

// this method walks through all chess and check if current chess is valid
// it can be time-consuming
// a better way is to check if one a put a new queen in position (row,col)
// see below
func isvalid(chess []string) bool {
	for row := range len(chess) {
		for col := range len(chess[0]) {
			if chess[row][col] != 'Q' {
				continue
			}
			// get Q position
			for offset := 1; offset < len(chess)-row; offset++ {
				if chess[row+offset][col] == 'Q' {
					return false
				}
				if col+offset < len(chess[0]) && chess[row+offset][col+offset] == 'Q' {
					return false
				}
				if col-offset >= 0 && chess[row+offset][col-offset] == 'Q' {
					return false
				}
			}
		}
	}
	return true
}
```



more efficient way of checking

```go
func solveNQueens(n int) [][]string {

	res := make([][]string, 0)
	path := make([]string, 0)
	backtrack(n, path, &res)
	return res
}

func backtrack(n int, path []string, res *[][]string) {
	if len(path) == n {
		temp := make([]string, n)
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for choice := range n {
		if !isvalid(path, len(path), choice) {
			continue
		}

		newRow := make([]byte, n)
		for j := range n {
			newRow[j] = '.'
		}
		newRow[choice] = 'Q'
		path = append(path, string(newRow))

		backtrack(n, path, res)

		path = path[:len(path)-1]
	}
}

// if you can put a queen in position row, col
func isvalid(chess []string, row int, col int) bool {

	// up
	for i := range row {
		if chess[i][col] == 'Q' {
			return false
		}
	}

	// up left and up right
	for offset := 1; offset <= row; offset++ {
		if col-offset >= 0 && chess[row-offset][col-offset] == 'Q' {
			return false
		}
		if col+offset < len(chess[0]) && chess[row-offset][col+offset] == 'Q' {
			return false
		}
	}

	return true
}
```





## 二进制枚举

枚举所有可能的二进制数

对于二进制枚举，比如n = 3, 枚举

000

001

010

011

100

101

110

111



### 子集， 子序列，子数组

对于没有重复的元素的数组，求子集和求子序列是同一个问题，且最终都有2^n个

```go
func subsets(nums []int) [][]int {
    n := len(nums)
    m := 1 << n
    
    ans := make([][]int, 0)
    
    for mask := 0; mask < m; mask++ {
        subset := make([]int, 0)
        for i := 0; i < n ; i++ {
            if mask & (1 << i) != 0{
                subset = append(subset, nums[i])
            }
        }
        ans = append(ans, subset)
        
    }
    return ans
    
}
```



对于递归回溯

```go

// 回溯解法
func subsets(nums []int) [][]int {
    
    ans := make([][]int, 0)
    subset := make([]int, 0)
    dfs(nums, 0, subset, &ans)
    return ans
}
func dfs(nums []int, i int, subset []int, ans *[][]int) {
    if i == len(nums) {
        *ans = append(*ans, append([]int{}, subset...))
        return
    }
    
    subset = append(subset, nums[i])
    dfs(nums, i+1, subset, ans)
    subset = subset[:len(subset)-1]
    dfs(nums, i+1, subset, ans)
    
}
```



而对于子数组而言，其必定是连续的，数目只有n * (n-1) / 2

```go
func subArray(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	for start := 0; start < n; start++ {
		for end := start; end < n; end++ {
			ans = append(ans, nums[start:end+1])
		}
	}
	return ans
}

```





### 组合

```go
func combine(n int, k int) [][]int {

	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)

	dfs = func(i int) {

		if len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return
		}
		if i > n {
			return
		}

		path = append(path, i)
		dfs(i + 1)
		path = path[:len(path)-1]
		dfs(i + 1)

	}
	dfs(1)
	return ans
}

// 一种更优化的剪枝
func combine(n int, k int) [][]int {

	ans := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(i int)

	dfs = func(i int) {

		if len(path)+(n-i+1) < k {
			return
		}

		if len(path) == k {
			ans = append(ans, append([]int{}, path...))
			return
		}

		path = append(path, i)
		dfs(i + 1)
		path = path[:len(path)-1]
		dfs(i + 1)

	}
	dfs(1)
	return ans
}
```





## 带有重复



### 全排列

nums: 1,1,2,

返回112,121,211,

```go
func permuteUnique(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	var dfs func(int)

	dfs = func(i int) {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}

		layer := make(map[int]bool)
		for i := range len(nums) {
			if layer[nums[i]] {
				continue
			}
			if used[i] {
				continue
			}
			layer[nums[i]] = true

			path = append(path, nums[i])
			used[i] = true

			dfs(i)

			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(0)
	return ans

}
```

力扣官方的解法：数组排序，然后相邻数判断

```go
func permuteUnique(nums []int) [][]int {

	sort.Ints(nums)
	ans := make([][]int, 0)
	path := make([]int, 0)
	used := make([]bool, len(nums))
	var dfs func(int)

	dfs = func(i int) {
		fmt.Println(path)
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := range len(nums) {
			if used[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			path = append(path, nums[i])
			used[i] = true

			dfs(i)

			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(0)
	return ans
}
```



### 子集

```
输入：nums = [1,2,2]
输出：[[],[1],[1,2],[1,2,2],[2],[2,2]]
```

```go
```

















