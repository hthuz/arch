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

