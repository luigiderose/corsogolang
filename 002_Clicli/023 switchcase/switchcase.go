package main

import (
	"fmt"
)

func main() {

	bond()
	james()

}

func bond() {
	switch {
	case false:
		fmt.Println("no stamp")
	case true:
		fmt.Println("si stamp")
		fallthrough
	case (3 == 2):
		fmt.Println("viene stampato anche se falso perchè c'è il falltrough")
	default:
		fmt.Println("non viene stampato")
	}
}

func james() { //caso multiplo
	x := "bond"
	switch x {
	case "nessuno", "bond", "cartepillar":
		fmt.Println("zio peppe")
	case "2":
		fmt.Println("non stampato")

	}
}
