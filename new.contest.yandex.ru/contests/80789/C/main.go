package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strconv"
)

type server struct {
	ttc int // time to completion
}

// serversPQ is priority queue of servers
type serversPQ []server

func (pq serversPQ) Len() int {
	return len(pq)
}

func (pq serversPQ) Less(i, j int) bool {
	return pq[i].ttc < pq[j].ttc
}

func (pq serversPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *serversPQ) Push(x any) {
	*pq = append(*pq, x.(server))
}

func (pq *serversPQ) Pop() any {
	old := *pq
	x := old[len(old)-1]
	*pq = old[:len(old)-1]
	return x
}

func main() {
	handle(os.Stdin, os.Stdout)
}

func newScanner(r io.Reader) *bufio.Scanner {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	return scanner
}

func scanInt(scanner *bufio.Scanner) int {
	if !scanner.Scan() {
		panic(io.ErrUnexpectedEOF)
	}
	num, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	return num
}

func handle(r io.Reader, w io.Writer) {
	scanner := newScanner(r)
	writer := bufio.NewWriter(w)
	defer writer.Flush()

	now := 0
	n := scanInt(scanner)
	k := scanInt(scanner)

	servers := make(serversPQ, k)
	heap.Init(&servers)

	for i := range n {
		t := scanInt(scanner)
		d := scanInt(scanner)

		s := heap.Pop(&servers).(server)
		now = max(t, s.ttc)
		s.ttc = now + d

		if i > 0 {
			fmt.Fprint(writer, " ")
		}

		fmt.Fprint(writer, s.ttc)

		heap.Push(&servers, s)
	}
}
