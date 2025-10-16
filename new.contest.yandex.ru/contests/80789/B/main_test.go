package main

import (
	"fmt"
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
			in: `10
1 5
2 7
1 3
4
3
2 8
2 9
3
4
4
`,
			out: `5 5
5 7
3 7
3 5
5 5
5 8
5 9
8 9
8 8
-1
`,
		},
		{
			in: `7
1 10
2 20
1 5
3
4
3
4
`,
			out: `10 10
10 20
5 20
10 20
10 10
-1
-1
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
		switch rand.IntN(4) + 1 {
		case pushFront:
			fmt.Fprintf(&input, "%d %d\n", pushFront, rand.IntN(1_000_000)+1)
		case pushBack:
			fmt.Fprintf(&input, "%d %d\n", pushBack, rand.IntN(1_000_000)+1)
		case popFront:
			fmt.Fprintln(&input, popFront)
		case popBack:
			fmt.Fprintln(&input, popBack)
		}
	}

	w := &strings.Builder{}

	for b.Loop() {
		handle(strings.NewReader(input.String()), w)
		w.Reset()
	}
}
