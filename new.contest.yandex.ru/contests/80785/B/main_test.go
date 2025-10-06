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
			in: `3
3 1 2 3
4 2 4 5 6
2 1 7
`,
			out: "7",
		},
		{
			in: `4
2 10 20
3 5 10 15
1 100
3 15 20 25
`,
			out: "6",
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
	input.Grow(10_000_000)
	input.WriteString("1000\n")
	for range 1000 {
		input.WriteString("1000")
		for range 1000 {
			input.WriteString(" " + strconv.Itoa(rand.IntN(1_000_000)+1))
		}
		input.WriteString("\n")
	}

	w := &strings.Builder{}

	for b.Loop() {
		handle(strings.NewReader(input.String()), w)
		w.Reset()
	}
}
