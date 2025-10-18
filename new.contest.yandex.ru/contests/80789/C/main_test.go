package main

import (
	"fmt"
	"io"
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
			in: `5 2
0 5
1 2
1 3
6 4
6 1
`,
			out: "5 3 6 10 7",
		},
		{
			in: `4 3
0 4
0 2
1 3
5 1
`,
			out: "4 2 4 6",
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

	input.WriteString("200000 200000\n")

	var t int

	for range 200000 {
		t += rand.IntN(5000)
		fmt.Fprintf(&input, "%d %d\n", t, rand.IntN(1_000_000_000)+1)
	}

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
