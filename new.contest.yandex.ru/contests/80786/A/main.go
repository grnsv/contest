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

const notFound int = -1

func main() {
	handle(os.Stdin, os.Stdout)
}

func handle(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	q, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}

	m := make(map[int]int, q/2)

	for range q {
		scanner.Scan()
		op, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		switch op {
		case add:
			scanner.Scan()
			x, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			scanner.Scan()
			y, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			m[x] = y
		case find:
			scanner.Scan()
			x, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}
			if y, ok := m[x]; ok {
				fmt.Fprintln(w, y)
			} else {
				fmt.Fprintln(w, notFound)
			}
		}
	}
}
