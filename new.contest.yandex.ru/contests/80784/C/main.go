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
	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())

	m := make(map[int]int, n)
	var x int

	for i := range n {
		scanner.Scan()
		x, _ = strconv.Atoi(scanner.Text())
		if _, ok := m[x]; !ok {
			m[x] = i
		}
	}

	for range q {
		scanner.Scan()
		x, _ = strconv.Atoi(scanner.Text())
		if i, ok := m[x]; ok {
			fmt.Fprintln(w, i+1)
		} else {
			fmt.Fprintln(w, -1)
		}
	}
}
