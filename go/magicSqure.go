package main

import "fmt"

func main() {
	grid := [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}
	fmt.Println(numMagicSquaresInside(grid))
}

func numMagicSquaresInside(grid [][]int) int {

	num := 0

	isMagic := func(row, col int) bool {
		occurs := make([]bool, 10)
		row_sum := make([]int, 3)
		col_sum := make([]int, 3)
		slope_sum := make([]int, 2)
		for r := row; r < len(grid) && r < row+3; r++ {
			for c := col; c < len(grid[0]) && c < col+3; c++ {
				if occurs[grid[r][c]] {
					return false
				}
				occurs[grid[r][c]] = true
				row_sum[r-row] += grid[r][c]
				col_sum[c-col] += grid[r][c]
				if (r-row == 0 && c-col == 0) || (r-row == 1 && c-col == 1) || (r-row == 2 && c-col == 2) {
					slope_sum[0] += grid[r][c]
				}
				if (r-row == 0 && c-col == 2) || (r-row == 1 && c-col == 1) || (r-row == 2 && c-col == 0) {
					slope_sum[1] += grid[r][c]
				}
			}
		}
		for _, sum := range row_sum {
			if sum != row_sum[0] {
				return false
			}
		}
		for _, sum := range col_sum {
			if sum != row_sum[0] {
				return false
			}
		}
		for _, sum := range slope_sum {
			if sum != row_sum[0] {
				return false
			}
		}
		return true
	}

	for row := 0; row < len(grid)-2; row++ {
		for col := 0; col < len(grid[0])-2; col++ {
			if isMagic(row, col) {
				num += 1
			}
		}
	}
	return num
}
