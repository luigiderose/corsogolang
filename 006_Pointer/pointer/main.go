package main

import "fmt"

var x int

func main() {
	x = 5
	Zero(x)
	fmt.Println(x)
	fmt.Printf("passa realmente dalla funzione Zero ma\nnon viene cambiato il valore in quanto\nnon agisce sulla parte della memoria\ndove viene salvato il valore della variabile non essendo un puntatore\n\n")
	PointerZero(&x)
	fmt.Println(x)
	fmt.Printf("Il puntatore agisce sulla memoria\ndove viene salvata la variabile e va a modificarne il valore\n\n")
	fmt.Println("il punto della memoria dove è salvato &x =", &x)
}

//Zero set fake zero
func Zero(x int) {
	x = 0
}

//PointerZero sert real zero
func PointerZero(xPtr *int) {
	*xPtr = 0
}
