package main

import (
	"fmt"
)

func main() {
	dichiarata()

	func() {
		fmt.Println("questa è una funzione anonima senza argomenti")
	}()

	func(x string) {
		fmt.Println(x)
	}("questa è una funzione non dichiarata con degli argomenti")

	fmt.Println("pippo")

}

func dichiarata() {
	fmt.Println("questa è una funzione dicharata con un nome")
}
