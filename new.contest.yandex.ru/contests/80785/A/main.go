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
	find
)

const (
	notFound int = iota
	found
)

func main() {
	handle(os.Stdin, os.Stdout)
}

func handle(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())
	set := make(map[int]struct{}, q/2)

	var (
		o int
		x int
	)

	for range q {
		scanner.Scan()
		o, _ = strconv.Atoi(scanner.Text())
		switch o {
		case add:
			scanner.Scan()
			x, _ = strconv.Atoi(scanner.Text())
			set[x] = struct{}{}
		case find:
			scanner.Scan()
			x, _ = strconv.Atoi(scanner.Text())
			if _, ok := set[x]; ok {
				fmt.Fprintln(w, found)
			} else {
				fmt.Fprintln(w, notFound)
			}
		}
	}
}
