package main

import (
	"fmt"
)

// esercizio crea una stringa e stampala
var x = 42
var y = "James Bond"
var z = true

func main() {

	s := fmt.Sprintf("%v - %v - %v", x, y, z)
	fmt.Println(s)
}
