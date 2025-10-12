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
	_, _ = fmt.Fscan(r, &n)
	ors := make(map[int]struct{})
	lastOrs := make(map[int]struct{})
	newOrs := make(map[int]struct{})

	for range n {
		for k := range newOrs {
			delete(newOrs, k)
		}

		var x int
		_, _ = fmt.Fscan(r, &x)

		newOrs[x] = struct{}{}
		ors[x] = struct{}{}

		for or := range lastOrs {
			newOr := or | x
			newOrs[newOr] = struct{}{}
			ors[newOr] = struct{}{}
		}

		lastOrs, newOrs = newOrs, lastOrs
	}

	fmt.Fprintln(w, len(ors))
}
