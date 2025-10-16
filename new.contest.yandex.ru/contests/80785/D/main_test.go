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
			in: `3
3 1 2 3
2 2 3
4 0 2 3 5
`,
			out: `YES
2
1 0 2
`,
		},
		{
			in: `3
2 1 2
2 2 3
2 1 3
`,
			out: `NO
`,
		},
		{
			in: `3
1 1
1 2
1 3
`,
			out: `YES
0
1 1 1
`,
		},
		{
			in: `3
3 1 2 4
3 2 3 4
3 1 3 4
`,
			out: `NO
`,
		},
		{
			in: `3
4 1 2 4 5
4 2 3 5 6
4 4 5 6 7
`,
			out: `NO
`,
		},
		{
			in: `3
0
0
0
`,
			out: `YES
0
0 0 0
`,
		},
		{
			in: `2
0
0
`,
			out: `YES
0
0 0
`,
		},
		{
			in: `2
2 1 2
2 2 3
`,
			out: `YES
1
1 1
`,
		},
		{
			in: `2
1 1
1 1
`,
			out: `YES
1
0 0
`,
		},
		{
			in: `2
1 1
2 1 2
`,
			out: `YES
1
0 1
`,
		},
		{
			in: `5
2 1 2
`,
			out: `YES
0
2 0 0 0 0
`,
		},
		{
			in: `3
`,
			out: `YES
0
0 0 0
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

	input.WriteString("400\n")
	for range 400 {
		input.WriteString("500")
		for range 500 {
			input.WriteString(" " + strconv.Itoa(rand.IntN(1_000_000_000)+1))
		}
		input.WriteString("\n")
	}

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
