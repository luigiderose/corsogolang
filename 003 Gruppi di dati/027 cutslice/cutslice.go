package main

import (
	"fmt"
)

func main() {

	/*make a slice*/
	x := []int{5, 7, 9, 5, 4, 7, 100, 15}
	/*print the slice*/
	fmt.Println(x)
	/*print the value in second position (base 0 = 1)*/
	fmt.Println(x[1])
	/*print all the slice from the second position at the end (use :)*/
	fmt.Println(x[1:])
	/*print the slice from the second position at 4 position (base 0 = 1:4 because it esclude the last value "use :")*/
	fmt.Println(x[0:4])

	/*for cicle with a range to print from 1 to 4 position*/

	for i := 1; i < 4; i++ {
		fmt.Println(i, x[i])
	}
}
