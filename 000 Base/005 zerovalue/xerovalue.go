package main

import (
	"fmt"
)

//dichiarare una variabile  e assegnare un valore a quella variabile
var y string
var z int

func main() {

	fmt.Println(y) //la variabile vale 0
	fmt.Printf("%T", y)

	y = "pietro pollo" //alla variabile viene assegnato un nuovo valore

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

}
