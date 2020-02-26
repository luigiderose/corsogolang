package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var x int
var Y *int
var count int
var second int
var Scelta string
var err error = nil

func main() {
	x = 1
	count = 0

	for count != 2 {

		fmt.Println(x)
		z := RealZero(&x)
		fmt.Println(z)
		fmt.Println(Y)
		fmt.Printf("il valore storato nle puntatore Y è %v, l'indirizo è %v\n", *Y, Y)
		z = RealSeven(&x)
		fmt.Printf("%v\n", z)
		fmt.Println(Y)
		fmt.Printf("il valore storato nle puntatore Y è %v, l'indirizo è %v\n", *Y, Y)
		fmt.Println(x)

		//reading a string
		reader := bufio.NewReader(os.Stdin) // inizializza il reader
		fmt.Println(reader)
		fmt.Println("metti il numero 1 se vuoi continuare")
		Scelta := "a" //err reader.ReadString('\n')// legget
		fmt.Println("Hai scelto ", Scelta)

		if err != nil {
			log.Println("errore")
		}

	}
	fmt.Println(count)
}

//FakeZero no change value
func FakeZero(y int) {
	y = 0

}

//RealZero change value
func RealZero(y *int) int {
	*y = 0
	Y = y

	return *Y

}

//RealSeven change value
func RealSeven(y *int) int {
	*y = 7
	Y = y

	return *Y
}
