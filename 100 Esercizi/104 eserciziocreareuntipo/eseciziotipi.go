package main

import (
	"fmt"
)

type azzeddotype string
type Luigi int

var x azzeddotype
var l Luigi
var y int

func main() {

	x = "azzeddomobili il mobile migliore per azzeddo"

	fmt.Println(x)
	fmt.Printf("\n")
	fmt.Println("è stato creato un nuovo tipo di variabile")
	fmt.Printf("\n")
	fmt.Printf("%T\n", x)

	fmt.Println(l)
	fmt.Printf("%T\n", l)
	l = 100
	fmt.Print(l)
	fmt.Printf(" , %T\n", l)

	y = int(l) + 1

	fmt.Print(y)
	fmt.Printf(" , %T\n", y)

	fmt.Scanln()

}
