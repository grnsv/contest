package main

import (
	"io"
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
			in: `5
3 2 5 4 1
`,
			out: `4 5
`,
		},
		{
			in: `7
6 2 5 1 7 3 4
`,
			out: `7 6
`,
		},
		{
			in: `3
0 1 2
`,
			out: `1 2
`,
		},
		{
			in: `1
-1000000000
`,
			out: `-1000000000
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

	input.WriteString("200000\n")

	for i := range 200_000 {
		if i > 0 {
			input.WriteRune(' ')
		}
		input.WriteString(strconv.Itoa(rand.IntN(2_000_000_001) - 1_000_000_000))
	}
	input.WriteRune('\n')

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
