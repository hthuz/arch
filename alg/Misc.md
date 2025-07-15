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
