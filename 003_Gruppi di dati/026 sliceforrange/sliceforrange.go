package main

import (
	"fmt"
)

func main() {
	/*create the slice*/
	x := []int{2, 7, 8, 9, 100, 1}
	/*assing at a variable the value of the lengh of the array*/
	y := len(x)
	/*these print the length of the slice*/
	fmt.Println(y)
	fmt.Println(len(x))
	/*these print the capacity of the slice*/
	fmt.Printf("%v\n\n\n", cap(x))

	/*these cycle print the content of the slice //pensato da me

	for i := 0; i < y; i++ {
		fmt.Println(x[i])
	}*/

	fmt.Println("contenuto dell'array slice con ciclo for")
	fmt.Printf("\n")
	/*più semplicemente//pensato da me stesso risultato di sopra metodo più semplice*/
	for v := range x {
		fmt.Println(v)
	}

	fmt.Printf("\n")
	/*these cycle print the position of any value and the value of the slice*/
	for z, k := range x {
		fmt.Println(z, k)
	}
}
