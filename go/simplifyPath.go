package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(simplifyPath("/../"))
	// s := "/home//user/Documents/../Pictures"
	s := "/a/../../b/../c//.//"
	// s := "/a/../../"
	fmt.Println(simplifyPath(s))

}

func simplifyPath(path string) string {
	sep := strings.Split(path, "/")
	stack := make([]string, 0)
	stack = append(stack, "")
	for _, c := range sep {
		if c == "." || c == "" {
			continue
		}
		if c == ".." {
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, c)
	}
	if len(stack) == 1 {
		return "/"
	} else {
		return strings.Join(stack, "/")
	}
}
