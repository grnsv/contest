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

	next := func() int {
		if !scanner.Scan() {
			panic(io.EOF)
		}
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		return n
	}

	n := next()
	numCounts := make(map[int]int, n)
	setSizes := make([]int, 0, n)

	for i := range n {
		k := next()
		for range k {
			x := next()
			numCounts[x]++
			count := numCounts[x]
			if count != 1 && count != i+1 {
				fail(w)
				return
			}
		}
		setSizes = append(setSizes, k)
	}

	centerSize := 0
	for _, count := range numCounts {
		if count == n {
			centerSize++
		} else if count != 1 {
			fail(w)
			return
		}
	}

	if centerSize == 0 {
		fail(w)
		return
	}

	fmt.Fprintf(w, "YES\n%d\n", centerSize)
	for i, size := range setSizes {
		if i == 0 {
			fmt.Fprint(w, size-centerSize)
		} else {
			fmt.Fprintf(w, " %d", size-centerSize)
		}
	}
}

func fail(w io.Writer) {
	fmt.Fprint(w, "NO")
}
