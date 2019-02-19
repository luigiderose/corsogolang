package main

import (
	"fmt"
)

func main() {
	x := []int{4, 16, 18, 32, 64}      //dichiaro una variabile slice
	fmt.Println(x)                     //la stampo
	x = append(x, 15, 16, 18, 19, 30)  //collego alla variabile che accoglie tutti i valori di x e aggiunge i nuovi valori
	fmt.Println(x)                     //la ristampo
	y := []int{40, 41, 42, 43, 44, 45} //dichiaro un'altra variabile slice
	x = append(x, y...)                //gli aggiungo tutta la y con i tre punti dopo vuol dire fino alla fine di y se non si mettono i punti errore non si può usare direttamente il cutting slice
	fmt.Println(x)                     //la ristampo
	z := y[1:3]                        //taglio la y e la assegno ad un'altra variabile z
	x = append(x, z...)                //aggiungo il taglio
	fmt.Println(x)                     //la ristampo

	/*con append possimo rimuovre componenti dello slice come da esempio sotto riportato dove eliminiamo la cifra 32 e 64 dallo slice x */
	x = append(x[:3], x[5:]...)
	fmt.Println(x)

}
