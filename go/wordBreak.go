package main

import (
	"fmt"
	"slices"
)

func main() {
	// s := "applepenapple"
	// dict := []string{"apple", "pen"}
	s := "catsandog"
	dict := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, dict))
}

func wordBreak(s string, wordDict []string) bool {
	canBreak := make([][]bool, len(s))
	for i, _ := range canBreak {
		canBreak[i] = make([]bool, len(s))
	}

	for start := len(s) - 1; start >= 0; start-- {
		for end := start; end < len(s); end++ {
			word := s[start : end+1]
			if slices.Contains(wordDict, word) {
				canBreak[start][end] = true
			} else {
				for i := 0; i < end-start; i++ {
					if canBreak[start][start+i] && canBreak[start+i+1][end] {
						canBreak[start][end] = true
						break
					}
				}
			}
		}
	}

	for _, row := range canBreak {
		fmt.Println(row)
	}

	return canBreak[0][len(s)-1]
}
