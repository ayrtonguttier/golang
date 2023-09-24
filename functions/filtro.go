package main

import "fmt"

type FiltroInteiro func(i int) bool

func main() {	
	var valores = []int{1, 2, 3, 4, 5}

	fmt.Println("Valores sem filtro: ", valores)

	valoresFiltrados := Filtrar(valores, func(i int) bool {
		return i > 2
	})

	fmt.Print("valores filtrados: ", valoresFiltrados)

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
