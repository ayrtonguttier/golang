package main

import "fmt"

type Numeravel interface {
	Numero() int
}

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

func (p Pessoa) Numero() int {
	return p.Idade
}

func swap[T Numeravel](numeros []T, a int, b int) {
	numeros[a], numeros[b] = numeros[b], numeros[a]
}

func BubbleSort[T Numeravel](numeros []T) {
	for i := 0; i < len(numeros)-1; i++ {

		if numeros[i].Numero() > numeros[i+1].Numero() {
			swap[T](numeros, i, i+1)			
		}
	}
}

func main() {
	pessoas := []Pessoa{{Nome: "Daniela", Sobrenome: "Guttier", Idade: 42}, {Nome: "Ayrton", Sobrenome: "Guttier", Idade: 29}}
	fmt.Println("Antes")
	for _, pessoa := range pessoas {
		fmt.Println(pessoa)
	}
	BubbleSort[Pessoa](pessoas)
	fmt.Println("Depois")
	for _, pessoa := range pessoas {
		fmt.Println(pessoa)
	}
}
