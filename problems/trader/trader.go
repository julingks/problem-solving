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
	maxArr := make([]int, n)
	max := arr[n-1]
	for i := n - 1; i >= 0; i-- {
		if max < arr[i] {
			max = arr[i]
		}
		maxArr[i] = max
	}
	fmt.Println(maxArr)
	profit := 0
	for i := 0; i < n; i++ {
		if maxArr[i]-arr[i] > profit {
			profit = maxArr[i] - arr[i]
		}
	}
	fmt.Println("answer:", profit)
}
