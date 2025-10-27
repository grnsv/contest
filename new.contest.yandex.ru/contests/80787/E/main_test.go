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
			in: `{[()]}
`,
			out: `YES
`,
		},
		{
			in: `([)]
`,
			out: `NO
`,
		},
		{
			in: `())(
`,
			out: `NO
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

	openBrackets := [3]byte{'(', '[', '{'}
	var stack [500000]byte

	for i := range stack {
		stack[i] = openBrackets[rand.IntN(len(openBrackets))]
		input.WriteByte(stack[i])
	}

	for i := len(stack) - 1; i >= 0; i-- {
		input.WriteByte(brackets[stack[i]])
	}

	fmt.Fprintln(&input)

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
