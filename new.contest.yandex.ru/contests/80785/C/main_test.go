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
			in: `3
3 1 2 3
4 2 3 4 5
2 2 3
`,
			out: "2",
		},
		{
			in: `3
4 1 2 3 4
3 2 3 5
5 2 3 4 6 7
`,
			out: "2",
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

	input.WriteString("1000\n")
	for range 1000 {
		input.WriteString("1000")
		for range 1000 {
			input.WriteString(" " + strconv.Itoa(rand.IntN(1_000_000)+1))
		}
		input.WriteString("\n")
	}

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
