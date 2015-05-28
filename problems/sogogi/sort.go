package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Play struct {
	Date      string
	MemberSeq int
}

type ByPlay []Play

func (b ByPlay) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByPlay) Less(i, j int) bool {
	if b[i].Date == b[j].Date {
		return b[i].MemberSeq < b[j].MemberSeq
	}
	return b[i].Date < b[j].Date
}

func (b ByPlay) Len() int {
	return len(b)
}

func main() {
	f, err := os.Open("input.csv")
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)

	arr := make([]Play, 0)
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
		seq, err := strconv.Atoi(strings.TrimSpace(memberSeq))
		if err != nil {
			panic(err)
		}
		arr = append(arr, Play{date, seq})
	}

	sort.Sort(ByPlay(arr))
	d := ""
	count := 0
	m := 0
	for _, v := range arr {
		if v.Date != d {
			if d != "" {
				fmt.Println(d, count)
			}
			d = v.Date
			m = v.MemberSeq
			count = 1
			continue
		}

		if v.MemberSeq != m {
			m = v.MemberSeq
			count++
		}
	}
	fmt.Println(d, count)
}
