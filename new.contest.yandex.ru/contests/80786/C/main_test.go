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
			in: `5
1 2
1 3
2 4
12 24
2 6
`,
			out: `1 2
`,
		},
		{
			in: `5
1 2
1 3
2 4
12 36
2 6
`,
			out: `1 3
`,
		},
		{
			in: `5
1 1
1 5
1 2
1 4
1 3
`,
			out: `1 5
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
	input.Grow(8_000)
	input.WriteString("1000\n")
	for range 1000 {
		input.WriteString(strconv.Itoa(rand.IntN(1000)+1) + " " + strconv.Itoa(rand.IntN(1000)+1) + "\n")
	}

	w := &strings.Builder{}

	for b.Loop() {
		handle(strings.NewReader(input.String()), w)
		w.Reset()
	}
}
