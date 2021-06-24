package main

import "fmt"

func main() {
	hello()

	arr := []float64{11, 27, 95, 89, 35}
	fmt.Println(average(arr))

	a, b := f()
	fmt.Printf("a = %v, b = %v \n", a, b)

	fmt.Println("total:", add(1, 2, 3))
}

// void function
func hello() {
	fmt.Println("Hello World")
}

// non void function
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

// multiple value function
func max(a, b int) int {
	
}

// multiple value return function
func f() (int, int) {
	return 1, 2
}

// variadic/variable arguments function
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
