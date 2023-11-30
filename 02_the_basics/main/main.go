package main

import "fmt"

type op func(int, int) int

func doIt(fn op, a, b int) int {
	return fn(a, b)
}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	sum := doIt(sum, 2, 3)
	sub := doIt(sub, 2, 3)

	fmt.Println(sum)
	fmt.Println(sub)
}
