package main

import (
	"fmt"
)

func main() {

	var numeriprimi [5]int //make an array with 5 position
	numeriprimi[0] = 2     //populate it
	numeriprimi[1] = 3
	numeriprimi[2] = 5
	numeriprimi[3] = 7
	numeriprimi[4] = 11

	for i, v := range numeriprimi { //utilizzo range per la stampa di ogni singolo valore per posizione
		fmt.Println(i, v)
	}

	fmt.Println(numeriprimi)        //print array
	fmt.Printf("%T\n", numeriprimi) //print tipe

}
