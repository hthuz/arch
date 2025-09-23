package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Header struct {
	Level int
	Title string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mdheadertree <markdown-file>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// First pass: collect all headers
	var headers []Header
	scanner := bufio.NewScanner(file)
	headerRe := regexp.MustCompile(`^(#+)\s*(.*)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := headerRe.FindStringSubmatch(line)
		if len(matches) == 3 {
			headers = append(headers, Header{
				Level: len(matches[1]),
				Title: strings.TrimSpace(matches[2]),
			})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Second pass: print with tree structure
	for i, h := range headers {
		// Determine if each level is the last in its group
		levelMarks := make([]bool, h.Level)
		for l := 0; l < h.Level; l++ {
			level := l + 1
			// Check if this is the last header at or above this level
			isLast := true
			for j := i + 1; j < len(headers); j++ {
				if headers[j].Level == level {
					isLast = false
					break
				}
				if headers[j].Level < level {
					break
				}
			}
			levelMarks[l] = isLast
		}

		// Build the tree prefix
		var prefix strings.Builder
		for l := 0; l < h.Level-1; l++ {
			if levelMarks[l] {
				prefix.WriteString("    ")
			} else {
				prefix.WriteString("│   ")
			}
		}

		if h.Level > 0 {
			if levelMarks[h.Level-1] {
				prefix.WriteString("└── ")
			} else {
				prefix.WriteString("├── ")
			}
		}

		fmt.Printf("%s%s\n", prefix.String(), h.Title)
	}
}
