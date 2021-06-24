package main

import "fmt"

func main() {
	r := factorial(20)
	fmt.Println(r)
}

func factorial(x int64) int64 {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
