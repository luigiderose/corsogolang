package main

import (
	"fmt"
)

func main() {

	if true { //vero stampa
		fmt.Println("001")
	}
	if false { //flaso non stampa
		fmt.Println("002")
	}

	if !true { //non è vero non stampa
		fmt.Println("003")
	}

	if !false { //non è flaso stampa
		fmt.Println("004")
	}

	if 2 == 2 { //vero stampa
		fmt.Println("005")
	}

	if 2 != 2 { //flaso non stampa
		fmt.Println("006")
	}

	if !(2 == 2) { //non è vero non stampa
		fmt.Println("007")
	}

	if !(2 != 2) { //non è flaso stampa
		fmt.Println("008")
	}

}
