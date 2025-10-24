package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	push int = iota + 1
	pop
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

func handle(r io.Reader, w io.Writer) {
	scanner := newScanner(r)

	q := scanInt(scanner)
	stack := make([]int, 0, q/2)

	for range q {
		switch scanInt(scanner) {
		case push:
			stack = append(stack, scanInt(scanner))
		case pop:
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}

		if len(stack) == 0 {
			fmt.Fprintln(w, -1)
		} else {
			fmt.Fprintln(w, stack[len(stack)-1])
		}
	}
}
