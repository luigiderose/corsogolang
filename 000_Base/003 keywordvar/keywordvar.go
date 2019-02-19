package main

import "fmt"

//dichiarazione lunga può essere fatta esternamente della funzione
var y = 101 //y ha un utilizzo per tutto il pacchetto
//dichiarazione di una variabile identificata "z" del tipo intero
var z int

func main() {

	//dichiarazione lampo
	x := 100 //la x ha lo scopo solo nella funzione main
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	x := 102
	fmt.Println("nuovamente", y)
	fmt.Println("la x è cambiata", x)
}
