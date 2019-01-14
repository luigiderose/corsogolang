package main

import (
	"fmt"
)

//dichiaro  costanti
const f string = "posso dichiarare una costante così" //in questo caso non ho la flessibilità type costant

//oppure così
const (
	a = 42 //in questo caso ho della flessibilita untype costant
	b = 42.78
	c = "James Bond"
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
