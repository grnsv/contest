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

type node struct {
	next, prev *node
	value      int
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

func handle(r io.Reader, w io.Writer) {
	scanner := newScanner(r)
	n := scanInt(scanner)

	var curr *node
	for range n {
		new := &node{value: scanInt(scanner)}
		if curr == nil {
			new.next, new.prev = new, new
			curr = new
			continue
		}
		new.prev = curr.prev
		new.next = curr
		new.prev.next = new
		new.next.prev = new
	}

	if n == 1 {
		fmt.Fprintln(w, curr.value)
		return
	}

	max := curr
	for n > 2 {
		curr = max
		prev, next := curr.prev, curr.next
		min := prev

		if next.value > max.value {
			max = next
		}
		if prev.value > max.value {
			max = prev
		}

		if curr.value < min.value {
			min = curr
		}
		if next.value < min.value {
			min = next
		}

		if curr == min {
			curr = min.next
		}

		min.prev.next = min.next
		min.next.prev = min.prev
		min.next = nil // avoid memory leaks
		min.prev = nil // avoid memory leaks
		n--
	}

	fmt.Fprintln(w, curr.value, curr.next.value)
}
