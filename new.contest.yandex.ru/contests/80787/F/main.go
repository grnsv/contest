package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	handle(os.Stdin, os.Stdout)
}

func newScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	return scanner
}

func scanInt(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		panic(io.ErrUnexpectedEOF)
	}
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func isOdd(n int) bool {
	return n%2 != 0
}

func handle(r io.Reader, w io.Writer) {
	scanner := newScanner(r)

	n := scanInt(scanner)
	stack := []int{-1}
	maxLen := 0

	for i := range n {
		switch isOdd(scanInt(scanner)) {
		case true:
			stack = append(stack, i)
		case false:
			if len(stack) > 1 {
				stack = stack[:len(stack)-1]
				maxLen = max(maxLen, i-stack[len(stack)-1])
			} else {
				stack[0] = i
			}
		}
	}

	fmt.Fprintln(w, maxLen)
}
