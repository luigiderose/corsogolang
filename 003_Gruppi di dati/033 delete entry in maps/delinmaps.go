package main

import (
	"fmt"
)

func main() {
	/*per cancellare un valore dalla maps utilizzare la seguente sintassi delete(<nomedellamaps>,"key")*/

	IDNomi := map[int]string{ //dichiaro una maps
		1: "Luigi",
		2: "Vanessa",
		3: "Luna",
		4: "Michelle",
	}

	fmt.Println(IDNomi) //la stampo

	delete(IDNomi, 1) //cancello il nome con chiave primaria = 1
	fmt.Println(IDNomi)

	if v, ok := IDNomi[4]; ok { //cerco se esiste ID Nome con valore 4 se non c'è esce
		fmt.Println("value:", v) //stampo il valore dell'idnome corrispondente a 4 ("Michelle")
		delete(IDNomi, 4)        //cancello il valore id 4
	}

	fmt.Println(IDNomi) //stampo il tutto

}
