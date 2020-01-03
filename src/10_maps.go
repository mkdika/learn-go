package main

import "fmt"

func main() {

	// map or hash declaration
	x := make(map[string]int)
	x["key"] = 10
	x["num"] = 55
	fmt.Println(x["key"])
	fmt.Println(x)
}
