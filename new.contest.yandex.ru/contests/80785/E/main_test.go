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
1 2 4
`,
			out: `6
`,
		},
		{
			in: `4
0 0 0 0
`,
			out: `1
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
	input.Grow(2_000_000)
	input.WriteString("200000\n")
	for i := range 200_000 {
		if i > 0 {
			input.WriteString(" ")
		}
		input.WriteString(strconv.Itoa(rand.IntN(1_000_000_001)))
	}

	w := &strings.Builder{}

	for b.Loop() {
		handle(strings.NewReader(input.String()), w)
		w.Reset()
	}
}
