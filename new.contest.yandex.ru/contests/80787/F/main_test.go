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
			in: `10
4 3 6 5 8 10 7 12 9 14
`,
			out: `4
`,
		},
		{
			in: `8
6 3 5 8 10 7 12 9
`,
			out: `6
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

	for range 200000 {
		input.WriteString(strconv.Itoa(rand.Int()))
		input.WriteByte(' ')
	}

	input.WriteByte('\n')

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
