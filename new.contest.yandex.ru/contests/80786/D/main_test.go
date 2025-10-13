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
			in: `6
1 2 3 2 1 1
`,
			out: `1 2 3
`,
		},
		{
			in: `6
3 1 1 2 3 3
`,
			out: `1 2 3
`,
		},
		{
			in: `8
11 21 31 41 31 41 11 21
`,
			out: `11 21 31
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
	input.Grow(600_000)
	input.WriteString("100000\n")
	for i := range 100_000 {
		if i != 0 {
			input.WriteString(" ")
		}

		input.WriteString(strconv.Itoa(rand.IntN(100_000) + 1))
	}

	w := &strings.Builder{}

	for b.Loop() {
		handle(strings.NewReader(input.String()), w)
		w.Reset()
	}
}
