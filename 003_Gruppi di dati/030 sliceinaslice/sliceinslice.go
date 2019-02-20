package main

import (
	"fmt"
)

func main() {

	x := []string{"Luigi", "è", "figo"} //dichiaro uno slice
	fmt.Println(x)
	y := []string{"zia", "peppa"} // dichiaro un secondo slice
	fmt.Println(y)
	xy := [][]string{x, y} //mlutidimensionale overo uno slice che contiene vari slice
	fmt.Println(xy)

	ContenitoreCifre := make([][]int, 6) //creo  uno slice di slice che determina quanti slice andrò a creare
	for i := 0; i < 6; i++ {             //qua creo con un ciclo for i 5 slice
		Lunghezza := i + 5                           //calcolo la lunghezza che avrà il primo slice in questo caso avrà 5 posizioni 0 + 5
		ContenitoreCifre[i] = make([]int, Lunghezza) //creo uno slice annidato nel primo richiamando lo slice di partenza e la variabile di riferimento
		for j := 0; j < Lunghezza; j++ {             //creo un ciclo for per poplare con i valori la j
			ContenitoreCifre[i][j] = (j + 1) * 2 //popolo lo slice j con i valori cerando così lo slice finale
		}
	}

	fmt.Println(ContenitoreCifre) //stampo
	fmt.Println(len(ContenitoreCifre))
	fmt.Println(cap(ContenitoreCifre))
}
