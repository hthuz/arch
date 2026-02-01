# Misc

Fastest way to know if two arrays contain the same element => using hash, $O(m+n)$

Fastest way to know if an array contain repetitive elements => using hash, $O(n)$



## Traverse directions

a more convenient to go up, down, left, right, four directions in 2-D array

```go 
type pair struct {r, c int}

var directions = []pair{{-1,0}, {1,0}, {0,-1}, {0,1}}

func traverse(row, col int) {
    for _, d := directions {
        newR, newC := row + d.r, col + d.c; 
        if 0 <= newR && newR < height && 0 <= newC && newC < width {
            // do something
        }
    }
}


```

## Deletion of an element from an array

By recording the updated length

```go
 
func del(nums []int, index int, length *int) {
    for i := index; i < length - 1; i++ {
        nums[i] = nums[i+1]
    }
    (*length)--
}
```



## isPrime

判断一个数是否为质数

```go

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i*i <= x; i++ {
		if (x/i)*i == x { /这里也可以等价为 if x % i == 0
			return false
		}
	}
	return true
}
```

因为(x/i)*i == x实际上是判断floor(x/i) == x/i

floor(x/i) = x/i 等价于x % i == -
