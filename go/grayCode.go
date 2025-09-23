package main

func main() {
	grayCode(4)
}

func grayCode(n int) []int {
	codes := []int{0, 1}
	next_codes := make([]int, 0)
	for range n - 1 {
		for _, c := range codes {
			next_codes = append(next_codes, c)
		}
		for i := len(codes) - 1; i >= 0; i-- {
			next_codes = append(next_codes, len(codes)+codes[i])
		}
		codes = next_codes
		next_codes = make([]int, 0)
	}
	return codes
}
