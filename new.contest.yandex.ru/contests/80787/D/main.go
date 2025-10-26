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

type deq struct {
	values []int
}

func (d *deq) isNotEmpty() bool {
	return len(d.values) != 0
}

func (d *deq) pushBack(v int) {
	d.values = append(d.values, v)
}

func (d *deq) front() int {
	return d.values[0]
}

func (d *deq) popFront() {
	d.values = d.values[1:]
}

func (d *deq) back() int {
	return d.values[len(d.values)-1]
}

func (d *deq) popBack() {
	d.values = d.values[:len(d.values)-1]
}

func handle(r io.Reader, w io.Writer) {
	scanner := newScanner(r)

	n := scanInt(scanner)
	k := scanInt(scanner)

	a := make([]int, n)
	for i := range a {
		a[i] = scanInt(scanner)
	}

	sum := 0
	minIndexes := &deq{}

	for i, v := range a {
		if minIndexes.isNotEmpty() && minIndexes.front() <= i-k {
			minIndexes.popFront()
		}

		for minIndexes.isNotEmpty() && a[minIndexes.back()] >= v {
			minIndexes.popBack()
		}

		minIndexes.pushBack(i)

		if i >= k-1 {
			sum += a[minIndexes.front()]
		}
	}

	fmt.Fprintln(w, sum)
}
