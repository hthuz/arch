# Backtrack

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



## Naive approach for full permutation:

Full permutation:
given [1,2,3], return [1,2,3]; [1,3,2]; [2,1,3]; [2,3,1]; [3,1,2]; [3,2,1]

```go
func permute(nums []int) [][]int {

	res := make([][]int, 0)
	if len(nums) == 1 {
		res = append(res, nums)
		return res
	}

	for i := range nums {
		pre_num := make([]int, i)
		copy(pre_num, nums[:i])
		post_num := nums[i+1:]

		sub_num := append(pre_num, post_num...)
		subperms := permute(sub_num)
		for _, subperm := range subperms {
			ans := append(subperm, nums[i])
			res = append(res, ans)
		}
	}

	return res
}
```



## backtrack approach for full permutation

```go
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	backtrack(nums, path, &res)
	return res
}

func backtrack(nums []int, path []int, res *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for _, num := range nums {
		// exclude illegal choices
		if slices.Contains(path, num) {
			continue
		}
		// make a choice, num here is a choice
		path = append(path, num)
		backtrack(nums, path, res)
		// undo choice
		path = path[:len(path)-1]
	}
}

```



or in a more straightforward form

```go
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	choices := make([]int, len(nums))
	copy(choices, nums)
	backtrack(choices, path, &res)
	return res
}

func backtrack(choices []int, path []int, res *[][]int) {
	if len(choices) == 0 {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for _, choice := range choices {
		path = append(path, choice)

		new_choices := make([]int, len(choices))
		copy(new_choices, choices)
		new_choices = slices.DeleteFunc(new_choices, func(e int) bool {
			return e == choice
		})

		backtrack(new_choices, path, res)

		path = path[:len(path)-1]
	}
}

```



Given N numbers

Number of nodes in the first layer 1

Number of nodes in the second layer N

Number of nodes in the third layer N(N-1)

Number of leaves:  N!

Time complexity: N!





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


## combineSum

Given `candidates`, each element unique, find all combinations that sum to `target` (leetcode 39)
you can use the same element for multiple times.

For example, candidates: [2,3,6,7], target: 7
The problem is that there may be repetition, for example, [2,2,3], [2,3,2], [3,2,2] so you need to check for repetition, which is inefficient

```go
func combinationSum(candidates []int, target int) [][]int {
    path := make([]int, 0)
    res := make([][]int, 0)

    var backtrack func(int, []int)
    backtrack = func(s int, path []int) {
        if s == target {
            slices.Sort(path)
            if is_sub(&path, &res) {
                return 
            }
            temp := make([]int, len(path))
            copy(temp, path)
            res = append(res, temp)

            return
        }
        if s > target {
            return
        }
        
        for _, c := range candidates {
            path = append(path, c)
            // note that you need to make a copy here, otherwise the answer is not correct
            temp := make([]int, len(path))
        	copy(temp, path)
			backtrack(s+c, temp)
            path = path[:len(path)-1]
        }
    }
    backtrack(0, path)
    return res
}

func is_sub(arr *[]int, arrs *[][]int) bool {
    for _, sub := range *arrs {
        same := true 
        if len(*arr) != len(sub) {
            same = false
            continue
        }
        for i := range len(*arr) {
            if (*arr)[i] != sub[i] {
                same = false 
                break
            }
        }
        if same {
            return true
        }
    }
    return false

}

```

A wiser way


```go
func combinationSum(candidates []int, target int) [][]int {

	path := make([]int, 0)
	res := make([][]int, 0)

	var backtrack func(int, int)
	backtrack = func(s int, index int) {

		if index >= len(candidates) {
			return
		}

		if s == target {
			res = append(res, append([]int{}, path...))
			return
		}

		if s > target {
			return
		}

		// move index forward
		backtrack(s, index+1)

		// keep the index
		path = append(path, candidates[index])
		backtrack(s+candidates[index], index)
		path = path[:len(path)-1]

	}
	backtrack(0, 0)

	return res
}

```



