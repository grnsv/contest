package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"testing"
)

func Test_handle(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			in: `5
1 5
2
1 6
1 7
2
`,
			out: `5
-1
6
6
7
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
	input.Grow(7_200_000)

	input.WriteString("1000000\n")

	for range 500_000 {
		fmt.Fprintf(&input, "1 %d\n", rand.IntN(1_000_000)+1)
	}
	for range 500_000 {
		switch rand.IntN(2) + 1 {
		case add:
			fmt.Fprintf(&input, "1 %d\n", rand.IntN(1_000_000)+1)
		case del:
			input.WriteString("2\n")
		}
	}

	w := &strings.Builder{}

	for b.Loop() {
		w.Grow(7_000_000)
		handle(strings.NewReader(input.String()), w)
		w.Reset()
	}
}
