package main

import (
	"fmt"
)

func main() {

	x := 1

	for { //ciclo for senza niente ciclo loop infinito

		if x > 10 {
			break
		}
		fmt.Println(x)
		x++

	}

	fmt.Println("done.")
}
