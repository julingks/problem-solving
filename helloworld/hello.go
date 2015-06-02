package main

import "fmt"

func main() {
	var cases int
	var name string
	fmt.Scanf("%d", &cases)
	for cases > 0 {
		fmt.Scanf("%s", &name)
		fmt.Printf("Hello, %s!\n", name)
		cases--
	}
}
