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

func handle(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	scanInt := func() int {
		if !scanner.Scan() {
			panic(io.ErrUnexpectedEOF)
		}
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		return num
	}

	n := scanInt()
	numCounts := make(map[int]int, n)

	for range n {
		numCounts[scanInt()]++
	}

	var result, resultCount int
	for x, count := range numCounts {
		if count < resultCount {
			continue
		}
		if count == resultCount && x > result {
			continue
		}

		result = x
		resultCount = count
	}

	fmt.Fprintln(w, result)
}
