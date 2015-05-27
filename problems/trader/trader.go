package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	var n int
	fmt.Fscanf(f, "%d", &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(f, "%d", &arr[i])
	}
	fmt.Println(arr)
}
