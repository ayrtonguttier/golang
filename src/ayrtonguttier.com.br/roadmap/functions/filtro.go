package functions

import (
	"fmt"

	"ayrtonguttier.com.br/roadmap/cli"
)

type FiltroInteiro func(i int) bool

func RunFilter() {	
	cli.Print("Filter")
	var valores = []int{1, 2, 3, 4, 5}

	fmt.Println("Valores sem filtro: ", valores)

	valoresFiltrados := Filtrar(valores, func(i int) bool {
		return i > 2
	})

	fmt.Println("valores filtrados: ", valoresFiltrados)	
}

func Filtrar(a []int, regra FiltroInteiro) []int {
	var resultado []int

	for _, val := range a {
		if regra(val) {
			resultado = append(resultado, val)
		}
	}

	return resultado
}
