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
5 2 3 4 1
`,
			out: "5 2 2 2 1",
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
