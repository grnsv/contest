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
			in: `9
2 5
1 5
2 5
1 6
1 10
2 7
1 7
2 10
2 7
`,
			out: `0
1
0
1
1
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

	input.WriteString("100000\n")
	for range 100000 {
		input.WriteString(strconv.Itoa(rand.IntN(2)+1) + " " + strconv.Itoa(rand.IntN(1_000_000_000)+1) + "\n")
	}

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
