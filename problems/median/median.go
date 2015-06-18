package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	r := bufio.NewReader(f)
	var n int
	fmt.Fscanf(r, "%d\n", &n)
	fmt.Println(n)

	l := list.New()

	for i := 0; i < n; i++ {
		var t int
		fmt.Fscanf(r, "%d", &t)
		l.PushBack(t)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	for {
		var command string
		fmt.Fscanf(r, "%s", &command)
		if command == "M" {
			fmt.Println("Median")
		} else if command == "R" {
			var idx int
			fmt.Fscanf(r, "%d", &idx)
			fmt.Println("Idx", idx)
		}
	}
}
