package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

type fraction struct {
	numerator   int
	denominator int
}

func newFraction(num, den int) fraction {
	div := gcd(num, den)
	return fraction{
		numerator:   num / div,
		denominator: den / div,
	}
}

func (f fraction) greaterThan(other fraction) bool {
	return f.numerator*other.denominator > other.numerator*f.denominator
}

func (f fraction) String() string {
	return fmt.Sprintf("%d %d", f.numerator, f.denominator)
}

func main() {
	handle(os.Stdin, os.Stdout)
}

func handle(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	scanInt := func() int {
		if !scanner.Scan() {
			panic(io.ErrUnexpectedEOF)
		}
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		return num
	}

	n := scanInt()
	fractions := make(map[fraction]int, n)

	for range n {
		fractions[newFraction(scanInt(), scanInt())]++
	}

	var (
		result      fraction
		resultCount int
	)
	for f, count := range fractions {
		if count < resultCount {
			continue
		}
		if count == resultCount && f.greaterThan(result) {
			continue
		}

		result = f
		resultCount = count
	}

	fmt.Fprintln(w, result)
}
