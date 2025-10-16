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
			in: `6
2 1 3 5 2 4
`,
			out: `2 4
4 5
`,
		},
		{
			in: `5
3 2 4 5 6
`,
			out: `2 5
1 2
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

	input.WriteString("100000\n")
	for range 100000 {
		input.WriteString(strconv.Itoa(rand.IntN(100_000)+1) + " ")
	}

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
