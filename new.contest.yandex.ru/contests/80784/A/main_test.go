package main

import (
	"math/rand/v2"
	"strconv"
	"strings"
	"testing"
)

func Test_handle(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in: `8
1 0 5
2 1
3 1
1 0 6
1 0 7
2 2
1 1 5
2 2
`,
			out: `5
6
5
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			w := &strings.Builder{}
			handle(strings.NewReader(tt.in), w)
			if got := w.String(); got != tt.out {
				t.Errorf("handle() = %v, want %v", got, tt.out)
			}
		})
	}
}

func Benchmark_handle(b *testing.B) {
	var input strings.Builder
	input.Grow(50000)
	input.WriteString("5000\n")
	for i := range 2000 {
		input.WriteString("1 " + strconv.Itoa(i) + " " + strconv.Itoa(rand.IntN(1_000_000)+1) + "\n")
	}
	for range 1000 {
		input.WriteString("2 " + strconv.Itoa(rand.IntN(2000)+1) + "\n")
	}
	for i := range 2000 {
		input.WriteString("3 " + strconv.Itoa(2000-i) + "\n")
	}

	w := &strings.Builder{}

	for b.Loop() {
		w.Grow(10000)
		handle(strings.NewReader(input.String()), w)
		w.Reset()
	}
}
