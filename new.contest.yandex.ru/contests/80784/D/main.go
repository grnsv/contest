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

	var curr, minimum int
	for i := range n {
		scanner.Scan()
		curr, _ = strconv.Atoi(scanner.Text())
		if i == 0 {
			minimum = curr
			fmt.Fprintf(w, "%d", minimum)
		} else {
			minimum = min(minimum, curr)
			fmt.Fprintf(w, " %d", minimum)
		}
	}
}
