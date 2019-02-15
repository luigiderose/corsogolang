package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano()) //genera attraverso il tempo attuale un seed unico in base alla data convertito in int64 tramite la funzione unixnano
	numero := randomInt(1, 101)      // a seconda del seed generato richiama la funazione randomint che calcola un numero randomico tra 1 e 99 numero minimo 1 massimo 100
	fmt.Println(numero)              // stampa il numero
	shiftdecrementabit(numero)       // chima la funzione decrementa bit di 1
	shiftincrementabit(numero)       // chima la funzione incrementa bit di 1
}

func shiftdecrementabit(a int) {
	b := a >> 1                          //decrementa a di 1 bit
	fmt.Printf("%d, %b, %#X\n", b, b, b) //stampa b in decimale binario e esadecimale
}

func shiftincrementabit(a int) {
	b := a << 1                          //incrementa a di 1 bit
	fmt.Printf("%d, %b, %#X\n", b, b, b) //stampa b in decimale binario e esadecimale
}

func randomInt(min, max int) int { //restituisce un intero maggiore uguale a min e minore di max
	return min + rand.Intn(max-min) //restituisce un valore random tra max e min
}
