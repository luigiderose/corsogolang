package main

import (
	"fmt"
	"luigiFunc/aggfunction"
)

func main() {

	fmt.Println("quanti anni hai?")
	anni := aggfunction.Eta()
	fmt.Printf("hai %v anni\n", anni)

}
