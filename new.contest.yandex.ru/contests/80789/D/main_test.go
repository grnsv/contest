package main

import (
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
			in: `6
ACDBCB
`,
			out: `ABCCBD
`,
		},
		{
			in: `4
DCBA
`,
			out: `ABCD
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
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for range 200_000 {
		input.WriteRune(letters[rand.IntN(len(letters))])
	}
	input.WriteString("\n")

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
