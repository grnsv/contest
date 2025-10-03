package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	handle(os.Stdin, os.Stdout)
}

func handle(r io.Reader, w io.Writer) {
	var n int
	fmt.Fscan(r, &n)

	a := make([]int, n)
	for i := range n {
		fmt.Fscan(r, &a[i])
	}

	minVal, minIdx := a[0], 0
	minDiff := 1000000
	minI, minJ := 0, 1

	maxVal, maxIdx := a[0], 0
	maxDiff := -1000000
	maxI, maxJ := 0, 1

	for j := 1; j < n; j++ {
		if diff := minVal - a[j]; diff < minDiff ||
			(diff == minDiff && (minIdx < minI || (minIdx == minI && j < minJ))) {
			minDiff = diff
			minI, minJ = minIdx, j
		}

		if diff := maxVal - a[j]; diff > maxDiff ||
			(diff == maxDiff && (maxIdx < maxI || (maxIdx == maxI && j < maxJ))) {
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
