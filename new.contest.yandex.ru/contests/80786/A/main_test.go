package main

import (
	"fmt"
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
1 2 3
2 1
2 2
1 1 4
2 1
2 2
1 2 5
2 1
2 2
`,
			out: `-1
3
4
3
4
5
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
	for range 100_000 {
		typ := rand.IntN(2) + 1
		switch typ {
		case 1:
			fmt.Fprintf(&input, "%d %d %d\n", typ, rand.IntN(1_000_000_000)+1, rand.IntN(1_000_000_000)+1)
		case 2:
			input.WriteString(strconv.Itoa(typ) + " " + strconv.Itoa(rand.IntN(1_000_000_000)+1) + "\n")
		}
	}

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
