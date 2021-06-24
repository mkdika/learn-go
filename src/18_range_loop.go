package main

import "fmt"

// range loop is like for each
func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	for _, v := range arr {
		fmt.Println(v)
	}
}
