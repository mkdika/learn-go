package main

import "fmt"

// create new type Circle of struct
type Circle struct {
	x, y, r float64
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c)
	fmt.Println(c.x, c.y, c.r)
}
