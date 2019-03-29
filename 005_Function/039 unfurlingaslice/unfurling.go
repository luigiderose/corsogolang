package main

import (
	"fmt"
)

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	y, k := sum("james", xi...)
	fmt.Println(y, k)
}

func sum(z string, x ...int) (string, int) {

	sum := 0

	for _, v := range x {

		sum += v

	}
	return z, sum
}
