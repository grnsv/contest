package main

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	pushFront int = iota + 1
	pushBack
	popFront
	popBack
)

func main() {
	handle(os.Stdin, os.Stdout)
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
	writer := bufio.NewWriter(w)
	defer writer.Flush()

	l := list.New()

	for range scanInt(scanner) {
		op := scanInt(scanner)

		switch op {
		case pushFront:
			l.PushFront(scanInt(scanner))
		case pushBack:
			l.PushBack(scanInt(scanner))
		case popFront:
			if l.Len() != 0 {
				l.Remove(l.Front())
			}
		case popBack:
			if l.Len() != 0 {
				l.Remove(l.Back())
			}
		}

		if l.Len() == 0 {
			fmt.Fprintln(writer, -1)
		} else {
			fmt.Fprintf(writer, "%d %d\n", l.Front().Value, l.Back().Value)
		}
	}
}
