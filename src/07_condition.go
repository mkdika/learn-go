package main

import "fmt"

func main() {
	n := 51
	if n % 2 == 0 {
		fmt.Println("is even")
	} else if n % 2 == 1 {
		fmt.Println("is odd")
	}
}