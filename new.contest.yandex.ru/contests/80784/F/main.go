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
	var maxVal, maxIdx int
	for i := range n {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
		if a[i] >= maxVal {
			maxVal = a[i]
			maxIdx = i
		}
	}

	isFirst := true
	for i, num := range a {
		if i == maxIdx {
			continue
		}

		if isFirst {
			fmt.Fprintf(w, "%d", num)
			isFirst = false
		} else {
			fmt.Fprintf(w, " %d", num)
		}
	}
}
