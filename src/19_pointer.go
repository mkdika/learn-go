package main

import "fmt"

func main() {
	x := 5
	fmt.Println(x) // x is still 5
	toZero(&x)
	fmt.Println(x) // x is 0
}

func toZero(xPtr *int) {
	*xPtr = 0
}
