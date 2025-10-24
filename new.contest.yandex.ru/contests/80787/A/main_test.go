package main

import (
	"fmt"
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
			in: `5
1 5
2
1 6
1 7
2
`,
			out: `5
-1
6
7
6
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

	input.WriteString("1000000\n")

	for range 1_000_000 {
		switch rand.IntN(2) + 1 {
		case push:
			fmt.Fprintf(&input, "1 %d\n", rand.IntN(1_000_000)+1)
		case pop:
			input.WriteString("2\n")
		}
	}

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
