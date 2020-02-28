package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Utenti db
type Utenti struct {
	Utenti []Utente `json:"utenti"`
}

// Utente il nome utente e il nome
type Utente struct {
	ID     string `json:"ID"`
	Name   string `json:"Name"`
	Player Player `json:"player"`
}

// Player dati di gioco
type Player struct {
	Life int `json:"Life"`
	Att  int `json:"Att"`
	Dif  int `json:"Dif"`
}

// Handler è una funziona di prova
func Handler() {

	var Persons Utenti

	jeson, err := os.Open("utente.json")

	if err != nil {
		log.Fatal("error cannot open: ", err)
	}

	fmt.Println("Successfully Opened users.json")

	jsn, err := ioutil.ReadAll(jeson)
	if err != nil {
		log.Fatal("error read d.body: ", err)
	}

	err = json.Unmarshal(jsn, &Persons)
	if err != nil {
		log.Fatal("decoding error: ", err)
	}

	for i := 0; i < len(Persons.Utenti); i++ {
		fmt.Printf("ID: %v\nName: %v\nLife: %v\n\n", Persons.Utenti[i].ID, Persons.Utenti[i].Name, Persons.Utenti[i].Player.Life)
	}
}

func main() {
	Handler()
}
