package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	add int = iota + 1
	del
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
	queue := make([]int, 0, q/2)

	for range q {
		op := scanInt(scanner)

		switch op {
		case add:
			x := scanInt(scanner)
			queue = append(queue, x)
		case del:
			queue = queue[1:]
		}

		if len(queue) == 0 {
			fmt.Fprintln(w, -1)
		} else {
			fmt.Fprintln(w, queue[0])
		}
	}
}
