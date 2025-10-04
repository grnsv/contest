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

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	a := make([]int, n)
	for i := range n {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	minVal, minIdx := a[0], 0
	minDiff := 1000000
	minI, minJ := 0, 1

	maxVal, maxIdx := a[0], 0
	maxDiff := -1000000
	maxI, maxJ := 0, 1

	var diff int
	for j := 1; j < n; j++ {
		diff = minVal - a[j]
		if diff < minDiff {
			minDiff = diff
			minI, minJ = minIdx, j
		}

		diff = maxVal - a[j]
		if diff > maxDiff {
			maxDiff = diff
			maxI, maxJ = maxIdx, j
		}

		if a[j] < minVal {
			minVal, minIdx = a[j], j
		}
		if a[j] > maxVal {
			maxVal, maxIdx = a[j], j
		}
	}

	fmt.Fprintf(w, "%d %d\n", minI+1, minJ+1)
	fmt.Fprintf(w, "%d %d\n", maxI+1, maxJ+1)
}
