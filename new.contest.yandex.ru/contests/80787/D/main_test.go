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
3
1 2 3 4 5
`,
			out: `6
`,
		},
		{
			in: `5
2
1 2 2 1 3
`,
			out: `5
`,
		},
		{
			in: `9
3
1 2 3 4 5 4 3 2 1
`,
			out: `16
`,
		},
		{
			in: `8
8
1 7 2 6 3 5 5 6
`,
			out: `1
`,
		},
		{
			in: `5
1
1 1 1 1 1
`,
			out: `5
`,
		},
		{
			in: `8
6
1 2 3 2 4 8 9 7
`,
			out: `5
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

	input.WriteString("300000\n100000\n")

	for range 300000 {
		fmt.Fprint(&input, rand.IntN(300000)+1, " ")
	}

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
