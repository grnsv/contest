package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type bytes [10]byte

func main() {
	handle(os.Stdin, os.Stdout)
}

func newScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	return scanner
}

func scanBytes(scanner *bufio.Scanner) []byte {
	if !scanner.Scan() {
		panic(io.ErrUnexpectedEOF)
	}
	return scanner.Bytes()
}

func scanInt(scanner *bufio.Scanner) int {
	num, err := strconv.Atoi(string(scanBytes(scanner)))
	if err != nil {
		panic(err)
	}
	return num
}

func handle(r io.Reader, w io.Writer) {
	scanner := newScanner(r)

	n := scanInt(scanner)
	m := make(map[bytes]map[byte]uint32, n*10)
	result := 0

	for range n {
		word := scanBytes(scanner)
		var pattern bytes
		copy(pattern[:], word)

		for i, letter := range word {
			pattern[i] = '*'

			if _, ok := m[pattern]; !ok {
				m[pattern] = make(map[byte]uint32)
			}
			letters := m[pattern]
			for midLetter, count := range letters {
				if midLetter != letter {
					result += int(count)
				}
			}
			letters[letter]++

			pattern[i] = letter
		}
	}

	fmt.Fprintln(w, result)
}
