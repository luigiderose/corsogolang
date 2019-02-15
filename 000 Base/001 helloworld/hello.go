package main

import (
	"fmt"
)

//flow control:
// (1) sequence
// (2) loop; interative
// (3) conditional
/*
fmt " è il pacchetto "
.println "è la funzione del pacchetto"

i pacchetti sono consultabili su golang.org
nella lettura ...<qualcosa> parametro variabile
interface{} può rientrarci qualisasi tipo di valore che si vuole (strin, int, bool)

_ = butta via un risultato ottenuto vedi l'esempio sotto che da una parte butta via gli errori
nel secondo caso segna che il valore di byte degli errori e nullo

*/

func main() {

	n, _ := fmt.Println("hello world, i'm new", 42, true)
	fmt.Println(n)

	z, err := fmt.Println("hello world a, i'm new", 42, true)
	fmt.Println(z)
	fmt.Println(err)

	foo()

	fmt.Println("after fooooooooooooooooooooooooooooo")

	for i := 0; i <= 100; i++ {

		if i%2 == 0 {

			fmt.Println(i)
		}

	}

	exit()

}

func foo() {

	fmt.Println("fooooooooooooooooooooooooooo")

}

func exit() {

	fmt.Println("we are exit")
}
