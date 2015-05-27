package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bs, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(bs), "\n")
	jump, err := strconv.Atoi(lines[0])
	if err != nil {
		panic("error")
	}
	n, err := strconv.Atoi(lines[1])
	if err != nil {
		panic("error")
	}
	pos := strings.Split(lines[2], " ")
	fmt.Println("jump:", jump)
	fmt.Println("n:", n)
	fmt.Println("positions:", pos)

	num := make([]int, 0)
	for _, v := range pos {
		p, err := strconv.Atoi(v)
		if err != nil {
			panic("error")
		}
		num = append(num, p)
	}

	first := num[0]
	under := make([]int, 0)
	over := make([]int, 0)

	for _, v := range num {
		if v < first {
			under = append(under, v)
		} else if v > first {
			over = append(over, v)
		}
	}

	//sort.Ints(under)
	sort.Sort(sort.Reverse(sort.IntSlice(under)))
	sort.Ints(over)

	fmt.Println("under :", under)
	fmt.Println("first :", first)
	fmt.Println("over :", over)

	can := 1
	prev := first
	for _, cur := range under {
		if prev-cur <= jump {
			can++
			prev = cur
		}
		break
	}

	prev = first
	for _, cur := range over {
		if cur-prev <= jump {
			can++
			prev = cur
		}
		break
	}
	fmt.Println("outout:", can)
}
