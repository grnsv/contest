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
	set := make(map[int]struct{}, n)
	var k, x int
	for range n {
		scanner.Scan()
		k, _ = strconv.Atoi(scanner.Text())
		for range k {
			scanner.Scan()
			x, _ = strconv.Atoi(scanner.Text())
			set[x] = struct{}{}
		}
	}

	fmt.Fprint(w, len(set))
}
