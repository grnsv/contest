package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	handle(os.Stdin, os.Stdout)
}

// hack: use array instead of map
var brackets = [124]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

func handle(r io.Reader, w io.Writer) {
	s, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	stack := make([]byte, 0, len(s)/2)
	isValid := true

	for _, b := range s {
		if b == '\n' {
			break
		}
		switch b {
		case '(', '[', '{':
			stack = append(stack, b)
		case ')', ']', '}':
			if len(stack) == 0 || brackets[stack[len(stack)-1]] != b {
				isValid = false
				break
			}
			stack = stack[:len(stack)-1]
		}
	}

	if isValid && len(stack) == 0 {
		fmt.Fprintln(w, "YES")
	} else {
		fmt.Fprintln(w, "NO")
	}
}
