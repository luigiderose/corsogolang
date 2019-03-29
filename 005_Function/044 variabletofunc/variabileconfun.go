package main

import (
	"fmt"
	"time"
)

func main() {

	secondi := func(x int) {
		fmt.Println("15.26.", x)
		time.Sleep(1 * time.Second)
	}

	for a := 0; a < 60; a++ {
		secondi(a)
	}

}
