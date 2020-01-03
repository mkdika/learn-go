package main

import "fmt"

func main() {
	// explicit variable declaration
	var str string = "hello world 3rd"
	fmt.Println(str)

	// implicit variable declaration with type inference
	s := "hello world 4th"
	fmt.Println(s)
}
