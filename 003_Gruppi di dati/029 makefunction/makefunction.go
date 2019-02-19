package main

import (
	"fmt"
)

func main() {
	//il comando make permette di gestire la creazione degli array gestendo la quantità di memoria da allocare in maniera intelligente
	x := make([]int, 10, 12) //creiamo un'array intelligente di dimensione minima di 10 posizioni massima 12
	fmt.Println(x)
	fmt.Println(len(x)) //stampiamo lunghezza
	fmt.Println(cap(x)) //stampiamo capienza
	x[0] = 1            //allochiamo in prima posizione il valore 1
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[9] = 9875 //allochiamo in ultima posizione il valore 9875
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 15) //aggiungiamo alla posizione successiva alla 10 il numero 15 avendo a disposizione 2 posizioni l'array aumenterà la lunghezza ma non la capienza
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 16)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 17) //dopo aver superato la capienza disponibile automaticamente la capienza viene raddoppiata allocando la memoria con un altro array
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
