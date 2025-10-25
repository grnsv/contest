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

type node struct {
	value, count int
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

	n := scanInt(scanner)
	a := make([]node, n)

	for i := range n {
		a[i] = node{
			value: scanInt(scanner),
			count: 0,
		}
		if i == 0 {
			fmt.Fprint(w, a[i].count)
			continue
		}

		z := i - 1
		for z > -1 && a[z].value < a[i].value {
			a[i].count += 1 + a[z].count
			z -= 1 + a[z].count
		}

		fmt.Fprint(w, " ", a[i].count)
	}
	fmt.Fprintln(w)
}
