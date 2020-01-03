package main

import "fmt"

func main() {

	// multiple variable declaration
	var (
		a int = 1
		b int = 2
		c int = 3
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// even for const
	const (
		x int = 4
		y int = 5
		z int = 6
	)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
