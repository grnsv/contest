package main

import (
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
rom
bom
dom
bot
rot
`,
			out: `6
`,
		},
		{
			in: `3
aa
aa
aa
`,
			out: `0
`,
		},
		{
			in: `6
aaa
aaB
aBa
Baa
BBB
abb
`,
			out: `3
`,
		},
		{
			in: `3
aa
aa
ab
`,
			out: `2
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
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	for range 100_000 {
		for range 10 {
			input.WriteRune(letters[rand.IntN(len(letters))])
		}
		input.WriteString("\n")
	}

	for b.Loop() {
		handle(strings.NewReader(input.String()), io.Discard)
	}
}
