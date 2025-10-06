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
			in: `5 4
4 2 1 5 2
1
2
4
6
`,
			out: `3
2
1
-1
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
	input.Grow(600000)
	input.WriteString("100000 100000\n")
	for range 100000 {
		input.WriteString(strconv.Itoa(rand.IntN(100_000)+1) + " ")
	}
	for range 100000 {
		input.WriteString(strconv.Itoa(rand.IntN(100_000)+1) + "\n")
	}

	w := &strings.Builder{}

	for b.Loop() {
		w.Grow(500_000)
		handle(strings.NewReader(input.String()), w)
		w.Reset()
	}
}
