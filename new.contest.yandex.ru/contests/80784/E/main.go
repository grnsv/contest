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

	result := make([]int, 0, n)
	lastIndex := len(a) - 1
	for i, num := range a {
		if i > 0 && i < lastIndex && a[i-1] > num && num < a[i+1] {
			continue
		}
		result = append(result, num)
	}

	fmt.Fprintf(w, "%d\n", len(result))
	for i, num := range result {
		if i == 0 {
			fmt.Fprintf(w, "%d", num)
		} else {
			fmt.Fprintf(w, " %d", num)
		}
	}
}
