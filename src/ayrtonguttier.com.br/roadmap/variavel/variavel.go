package variavel

import (
	"fmt"
)

func RunVariavel(){

	//declaração explicita
	var nome string = "Ayrton"
	//declaração explicita com inferencia de tipo
	var sobrenome = "Guttier"
	//declaração inferencia de tipo
	idade:= 29
	
	fmt.Println(nome, sobrenome, idade)
}
