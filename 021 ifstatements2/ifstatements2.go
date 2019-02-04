package main

import (
	"fmt"
)

func main() {
	//dichiarazione rapida in un if

	if x := 42; x == 42 {
		fmt.Println("ok")
	}
	/*fmt.println(x) provo a richiamare la x fuoir dalle parentesi quadre mi darà errore in quanto la x è dichiarata all'interno dello statements if e non è globale la dovrei ridichiarare
	per farla funzionare dovrei dichiararla prima dell'if

	esempio
	*/
	x := 42
	if x == 42 {
		fmt.Println("ok")
	}
	fmt.Println(x)
	/*mi stamperebbe la x */
}
