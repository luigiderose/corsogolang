package main

import (
	"fmt"
)

func main() {
	//dichiaro una variabile
	x := 4
	fmt.Printf("%d, decimale\t\t%b, binario\n", x, x) //cosa ho scritto stampa il primo valore decimale poi due tab e il numero in binario dopo la virgola la variabile dichiarata in precedenza

	y := x << 1 //shifto il bit di x di uno verso sinistra (lo incremento) creando il numero 8 decimale che corrisponde al passaggio da 100 a 1000 binario
	fmt.Printf("%d, decimale\t\t%b, binario\n", y, y)

	z := y >> 2
	fmt.Printf("%d, decimale\t\t%b, binario\n", z, z) //shifto di 2 posizioni il bit verso destra (lo decremento) creando il numero 2 decimale che corrisponde al passaggio da 1000 a 10 binario
}
