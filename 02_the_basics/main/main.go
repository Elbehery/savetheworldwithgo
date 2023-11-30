package main

import "fmt"

func accum(delta int) func() int {
	i := 0
	return func() int {
		i += delta
		return i
	}
}

func main() {
	byOne := accum(1)
	byTwo := accum(2)

	for i := 0; i < 5; i++ {
		fmt.Println(byTwo())
		fmt.Println(byOne())
		fmt.Printf("-------\n")
	}
}
