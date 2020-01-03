package main

import "fmt"

func main() {
	// normal variable can be change later
	var s string = "normal variable"
	fmt.Println(s)
	s = "can be changed after declaration"
	fmt.Println(s)

	// but constants variable is final after declaration
	const str string = "This is constants variable"
	fmt.Println(str)

	// change const variable will cause compile time error
	// str = "I like to change ;)"
}
