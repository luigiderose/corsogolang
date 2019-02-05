package main

import (
	"fmt"
)

func main() {

	Esercizio1()
	Esercizio2()
	Esercizio3()
	Esercizio4()
	Esercizio5()
	Esercizio6()
	Esercizio7()
	Esercizio8()
	Esercizio9()
	Esercizio10()
}

func Esercizio1() {

	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}

}

func Esercizio2() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for z := 1; z <= 3; z++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func Esercizio3() {
	for agestart := 1986; agestart <= 2019; agestart++ {
		fmt.Println(agestart)
	}
	fmt.Println("")

}

func Esercizio4() {
	agestart := 1986
	for {
		if agestart > 2019 {
			break
		}
		fmt.Println(agestart)
		agestart++
	}
	fmt.Println("")
}

func Esercizio5() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("se %v è diviso per quattro, il resto (o modulus) è %v\n", i, i%4)
	}
	fmt.Println("")
}

func Esercizio6() {
	x := "pippo"
	if x == "pippo" {
		fmt.Println("stica")
	}
	fmt.Println("")
}

func Esercizio7() {
	x := "pippo"
	if x == "crauti" {
		fmt.Println("stica")
	} else if x == "pippo" {
		fmt.Println("suocera")
	} else {
		fmt.Println("nochanche")
	}
	fmt.Println("")
}

func Esercizio8() {

	switch {
	case !true:
		fmt.Println("vero")
	case !false:
		fmt.Println("flase")
	}
	fmt.Println("")
}

func Esercizio9() {
	favSport := "pallavolo"
	switch favSport {
	case "calcio":
		fmt.Println("calcio")
	case "pallavolo":
		fmt.Println("pallavolo")
	}
	fmt.Println("")
}

func Esercizio10() {
	fmt.Println(true && true)  // true And
	fmt.Println(true && false) // false And
	fmt.Println(true || true)  // true
	fmt.Println(true || false) // true
	fmt.Println(!true)         // flase
}
