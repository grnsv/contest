package main

import (
	"fmt"
	"io"
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
1 2 3 4 5
`,
			out: `0 1 2 3 4
`,
		},
		{
			in: `5
1 2 2 1 3
`,
			out: `0 1 0 0 4
`,
		},
		{
			in: `9
1 2 3 4 5 4 3 2 1
`,
			out: `0 1 2 3 4 0 0 0 0
`,
		},
		{
			in: `8
1 7 2 6 3 5 5 6
`,
			out: `0 1 0 1 0 1 0 3
`,
		},
		{
			in: `5
1 1 1 1 1
`,
			out: `0 0 0 0 0
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

	input.WriteString("300000\n")

	for i := range 300000 {
		fmt.Fprint(&input, i+1, " ")
	}

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
