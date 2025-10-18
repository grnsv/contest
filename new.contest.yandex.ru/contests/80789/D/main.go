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
	writer := bufio.NewWriter(w)
	defer writer.Flush()

	var n int
	if _, err := fmt.Fscanln(r, &n); err != nil {
		panic(err)
	}

	s := make([]byte, n)
	if _, err := r.Read(s); err != nil {
		panic(err)
	}

	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] <= s[j] {
			if err := writer.WriteByte(s[i]); err != nil {
				panic(err)
			}
			i++
		} else {
			if err := writer.WriteByte(s[j]); err != nil {
				panic(err)
			}
			j--
		}
	}

	if _, err := writer.WriteRune('\n'); err != nil {
		panic(err)
	}
}
