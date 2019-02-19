package main

import (
	"fmt"
)

func main() {

	//devo sampare solo i numeri dispari da 0 a 100

	x := 0 //variabile contatore

	for { //loop infinito

		if x > 100 {
			break
		}

		if x%2 == 0 {
			fmt.Println(x)
		}
		x++

	}

	for y := 0; y <= 100; y++ {
		if y%2 != 0 {
			fmt.Println(y)
		}
	}

}
