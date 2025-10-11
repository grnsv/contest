package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	handle(os.Stdin, os.Stdout)
}

func handle(r io.Reader, w io.Writer) {
	in := bufio.NewReader(r)
	out := bufio.NewWriter(w)
	defer out.Flush()

	var n int
	_, _ = fmt.Fscan(in, &n)

	numCounts := make(map[int]int, n)
	setSizes := make([]int, 0, n)

	for i := range n {
		var k int
		_, _ = fmt.Fscan(in, &k)
		set := make(map[int]struct{}, k)
		for range k {
			var x int
			_, _ = fmt.Fscan(in, &x)
			if _, ok := set[x]; ok {
				continue
			}
			set[x] = struct{}{}
			numCounts[x]++
			count := numCounts[x]
			if count != 1 && count != i+1 {
				fail(out)
				return
			}
		}
		k = len(set)
		setSizes = append(setSizes, k)
	}

	centerSize := 0
	for _, count := range numCounts {
		if count == n {
			centerSize++
		} else if count != 1 {
			fail(out)
			return
		}
	}

	fmt.Fprintf(out, "YES\n%d\n", centerSize)
	for i, size := range setSizes {
		if i == 0 {
			fmt.Fprint(out, size-centerSize)
		} else {
			fmt.Fprintf(out, " %d", size-centerSize)
		}
	}
	fmt.Fprintln(out)
}

func fail(w io.Writer) {
	fmt.Fprintln(w, "NO")
}
