package main

import (
	"fmt"
)

func main() {
	x := 43

	if x == 40 {
		fmt.Println("il nostro numero è 40")
	} else if x == 41 { //va tutto sulla stessa linea
		fmt.Println("il nostro numero è 41")
	} else if x == 42 { //va tutto sulla stessa linea
		fmt.Println("il nostro numero è 42")
	} else if x == 43 { //va tutto sulla stessa linea
		fmt.Println("il nostro numero è 43")
	} else { //va tutto sulla stessa linea
		fmt.Println("il nostro numero non è 40, 41, 42 o 43 ")
	}

}
