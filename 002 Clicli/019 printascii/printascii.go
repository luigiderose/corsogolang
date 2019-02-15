package main

import (
	"fmt"
)

func main() {

	for i := 33; i <= 122; i++ {
		fmt.Printf("%c\t%v\t%#x\t%#U\n", i, i, i, i)
	}

}
