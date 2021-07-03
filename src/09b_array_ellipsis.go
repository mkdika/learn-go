package main

import "fmt"

func main() {
	// use an ellipsis (...) to use an implicit length when you
	// pass the values
	a := [...]string{"hello", "world"}
	fmt.Printf("%q", a)
}
