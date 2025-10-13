package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
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

type pair struct {
	x     int
	count int
}

func handle(r io.Reader, w io.Writer) {
	scanner := newScanner(r)

	n := scanInt(scanner)
	m := make(map[int]int, n)
	for range n {
		m[scanInt(scanner)]++
	}

	pairs := make([]pair, 0, len(m))
	for x, count := range m {
		pairs = append(pairs, pair{x: x, count: count})
	}

	slices.SortFunc(pairs, func(a, b pair) int {
		if a.count != b.count {
			return b.count - a.count
		}
		return a.x - b.x
	})

	result := []int{pairs[0].x, pairs[1].x, pairs[2].x}
	slices.Sort(result)

	fmt.Fprintf(w, "%d %d %d\n", result[0], result[1], result[2])
}
