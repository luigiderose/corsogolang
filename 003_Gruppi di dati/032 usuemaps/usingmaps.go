package main

import "fmt"

func main() {

	IDNomi := map[int]string{ //dichiaro la variabile maps e la popolo
		0: "Luigi",
		1: "Vanessa",
		2: "Luna",
		3: "Michelle",
	}

	for i, v := range IDNomi { // dichiaro due variabili una per il contegggio la seconda per il valore da restituire con il comando range + nome variabile della maps stampo tutti i valori
		if i%2 == 0 {
			fmt.Println(i, v)
		}
	}

}
