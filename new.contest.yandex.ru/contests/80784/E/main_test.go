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
1 3 2 5 4`,
			out: `4
1 3 5 4`,
		},
		{
			in: `5
3 2 1 2 3`,
			out: `4
3 2 2 3`,
		},
		{
			in: `6
5 3 1 2 4 6`,
			out: `5
5 3 2 4 6`,
		},
		{
			in: `7
2 1 2 1 2 1 2`,
			out: `4
2 2 2 2`,
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
	input.Grow(600000)
	input.WriteString("100000\n")
	for range 100000 {
		input.WriteString(strconv.Itoa(rand.IntN(99999)+1) + " ")
	}

	w := &strings.Builder{}
	w.Grow(400000)

	for b.Loop() {
		handle(strings.NewReader(input.String()), w)
		w.Reset()
	}
}
