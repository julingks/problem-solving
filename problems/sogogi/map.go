package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.csv")
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)

	m := make(map[string]map[string]int)
	for {
		date, err := r.ReadString(',')
		if err == io.EOF {
			break
		}
		date = date[:10]
		if err != nil {
			panic(err)
		}
		memberSeq, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		memberSeq = memberSeq[:len(memberSeq)-1]

		if _, ok := m[date]; !ok {
			m[date] = make(map[string]int)
		}
		m[date][memberSeq] = 1
	}

	for k, v := range m {
		fmt.Println(k, len(v))
	}
}
