package aggfunction

import "fmt"

//MediumLife vita media
const (
	MediumLife int = 150
	pi             = 3.14
)

var (
	age  int
	name string
)

//Eta inizializza l'eta
func Eta() int {

	_, err := fmt.Scan(&age)
	CotrolloErrori(err)
	ValidaEta(age)
	return age
}

//ValidaEta valida l'età se sopra i 150 anni chiede di reinserire
func ValidaEta(z int) {

	if z >= MediumLife {

		fmt.Println("Età sopra l'età media?")
		//aggiungere scan per si o no

		_, err := fmt.Scan(&age)

		CotrolloErrori(err)
	}
}

//CotrolloErrori controllo l'errore
func CotrolloErrori(err error) {

	for err != nil {
		selecterror := err.Error()

		switch selecterror {

		case "expected integer":

			fmt.Println("errore sull'età, inserisci un un numero")
			_, err = fmt.Scan(&age)

		default:
			fmt.Println("errore 404")
		}

	}
}
