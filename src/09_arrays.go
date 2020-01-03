package main

import "fmt"

func main() {
	// array has fixed length
	var x [6]int
	x[4] = 100
	fmt.Println(x)

	nums := [5]int{1,2,3,4,5}
	fmt.Println(nums)
}