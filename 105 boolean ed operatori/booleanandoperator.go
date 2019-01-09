package main

import (
	"fmt"
)

var x bool // dichiaro una variabile booleana senza asssegnargli alcun valore

func main() {

	fmt.Println(x) // la stampo sarà falsa perchè vale 0

	x = true // gli assegno il valore vero

	fmt.Println(x) // otterrà il valore vero

	/******* utilizzo di alcuni operatori logici*******/

	//di asegnazione rapida   := assegna tipo e valore a una variabile

	a := 7  //in questo caso intero
	b := 42 //in questo caso intero

	fmt.Println(a >= b) //la risposta sarà falsa in quanto il risultato dell'operazione a maggiore uguale a b è falsa
	fmt.Println(a <= b) //la risposta sarà vero in quanto il risultato dell'operazione a minore uguale a b è vera
	fmt.Println(a != b) // a è diverso da b? si quindi true
	fmt.Println(a == b) // a è uguale da b? no quindi fale (== confronta) (= assegna)
}
