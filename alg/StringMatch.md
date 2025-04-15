# String Match

occurence of a pattern within a text string

```go
// find pattern in string
func naiveMatch(pat string, txt string) int {
	// note that i goes end to len(txt) - len(pat) instead of only len(txt)-1
	for i := 0; i <= len(txt)-len(pat); i++ {
		j := 0
		for j = 0; j < len(pat); j++ {
			if pat[j] != txt[i+j] {
				break
			}
		}
		if j == len(pat) {
			return i
		}
	}
	return -1
}
```

KMP from the perspective of DFA. The key lines in how you construct the DFA

```go 
func computeDFA(pat string) [][]int {

	// DFA[state][input] = next_state
	DFA := make([][]int, len(pat))
	for state := range DFA {
		DFA[state] = make([]int, 256)
	}

	DFA[0][pat[0]] = 1
	x := 0

	for j := 1; j < len(pat); j++ {
		for c := range 256 {
			if c == int(pat[j]) {
				DFA[j][c] = j + 1
			} else {
				DFA[j][c] = DFA[x][c]
			}
		}
		x = DFA[x][pat[j]]
	}

	return DFA
}

func KMP_DFA(pat string, txt string) int {

	DFA := computeDFA(pat)
	j := 0
	for i := range len(txt) {
		j = DFA[j][txt[i]]
		if j == len(pat) {
			return i - len(pat) + 1
		}
	}
	return -1
}

```

