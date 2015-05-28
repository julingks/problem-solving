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

	maxProfit := 0
	for i := 0; i < n; i++ {
		profit := 0
		for j := i + 1; j < n; j++ {
			if arr[j]-arr[i] > profit {
				profit = arr[j] - arr[i]
			}
		}
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	fmt.Println("answer : ", maxProfit)
}
