package main

import (
	"fmt"
)

func main() {

	/********* ciclo for classico *************/
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\t %#U\t %#X\n", i, i, i)
	}

	/********** ciclo for annidato v1 **********/
	for a := 1; a <= 10; a++ {
		for b := 1; b <= 3; b++ {
			fmt.Printf("%d\t%d\n", a, b)
		}
	}

	/********** ciclo for annidato v2 **********/
	for a := 1; a <= 10; a++ {
		b := 0
		fmt.Printf("%d\t%d\n", a, b)
		for b := 1; b <= 3; b++ {
			fmt.Printf("%d\t%d\n", a, b)
		}
	}

}
