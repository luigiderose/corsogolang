package main

import (
	"fmt"
)

type azzeddotype string

func main() {

	var x azzeddotype = "azzeddomobili il mobile migliore per azzeddo"

	fmt.Println(x)
	fmt.Printf("\n")
	fmt.Println("è stato creato un nuovo tipo di variabile")
	fmt.Printf("\n")
	fmt.Printf("%T\n", x)

	fmt.Scanln()

}
