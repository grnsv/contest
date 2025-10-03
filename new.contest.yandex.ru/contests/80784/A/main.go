package main

import (
	"fmt"
	"io"
	"os"
)

type operation uint8

const (
	add operation = iota + 1
	find
	remove
)

type linkedList struct {
	value uint32
	next  *linkedList
}

func (l *linkedList) add(value uint32, index uint16) *linkedList {
	var prev, next *linkedList
	next = l
	for range index {
		prev = next
		next = next.next
	}

	new := &linkedList{value: value, next: next}
	if prev == nil {
		l = new
	} else {
		prev.next = new
	}

	return l
}

func (l *linkedList) find(index uint16) uint32 {
	curr := l
	for range index - 1 {
		curr = curr.next
	}
	return curr.value
}

func (l *linkedList) remove(index uint16) *linkedList {
	var prev, curr *linkedList
	curr = l
	for range index - 1 {
		prev = curr
		curr = curr.next
	}
	if prev == nil {
		l = curr.next
	} else {
		prev.next = curr.next
	}
	curr = nil

	return l
}

func main() {
	handle(os.Stdin, os.Stdout)
}

func handle(r io.Reader, w io.Writer) {
	var count uint16
	_, _ = fmt.Fscan(r, &count)

	var lst *linkedList

	var (
		oper  operation
		index uint16
		value uint32
	)

	for range count {
		_, _ = fmt.Fscan(r, &oper)
		switch oper {
		case add:
			_, _ = fmt.Fscan(r, &index, &value)
			lst = lst.add(value, index)
		case find:
			_, _ = fmt.Fscan(r, &index)
			fmt.Fprintln(w, lst.find(index))
		case remove:
			_, _ = fmt.Fscan(r, &index)
			lst = lst.remove(index)
		}
	}
}
