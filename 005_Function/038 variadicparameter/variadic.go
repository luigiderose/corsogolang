package main

import (
	"fmt"
)

func main() {
	x := pippo(2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("azz", x)
}

func pippo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0

	for i, v := range x {
		sum += v
		fmt.Println("for item in index position, ", i, " we are now adding, ", v, " the total, is ", sum)
	}
	fmt.Println("the total is ", sum)
	return sum
}
