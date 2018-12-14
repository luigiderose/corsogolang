package main

import (
	"fmt"
)

func main() {

	init := "git init"
	Pull := "git pull https://github.com/luigiderose/corsogolang"
	AggiungiFile := "git add --all"
	Commit := "git commit -m"
	Push := "git push"

	fmt.Println("comandi git e descrizione")
	fmt.Print(init)
	fmt.Println(", da lanciare nella cartella per inizializzare la sincronizzazione git")
	fmt.Print(AggiungiFile)
	fmt.Println(", da lanciare nella cartella di reposity principale per aggiungere i file")
	fmt.Print(Commit)
	fmt.Println(", da lanciare nella cartella di reposity principale per eseguire il commit")
	fmt.Print(Push)
	fmt.Println(", da lanciare nella cartella di reposity principale per eseguire il push sul server")
	fmt.Print(Pull)
	fmt.Println(", da lanciare per eseguire la pull di un")

}
