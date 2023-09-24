package main

import (
	"fmt"
	"time"
)

func numeros() (int, int) {
	return 3, 3
}

func normal() {
	if 7%2 == 0 {
		fmt.Println("É par")
	} else {
		fmt.Println("É impar")
	}
}

func comInicializacao() {
	if r := 8 % 2; r == 0 {
		fmt.Println("É par")
	} else {
		fmt.Println("É impar")
	}
}

func comElseIf() {
	if a, b := numeros(); b < a {
		fmt.Println("b é menor que a")
	} else if b > a {
		fmt.Println("b é maior que a")
	} else {
		fmt.Println("b é igual a")
	}
}

func comBooleano() {
	x := true

	if x {
		fmt.Println("x é verdadeiro")
	}
}

func switchCase() {
	today := time.Now()
	switch today.Day() {
	case 23:
		fmt.Println("Dia vinte e três")
	default:
		fmt.Println("Outro dia")
	}

}

func switchCaseComVariavel() {

	today := time.Now()
	day := today.Day()

	switch day {
	case 23:
		fmt.Println("Dia vinte e três Com variavel")
	default:
		fmt.Println("Outro dia Com variavel")
	}

}

func switchCaseComInicializacao() {

	switch today := time.Now().Day(); today {
	case 23:
		fmt.Println("Dia vinte e três Com inicialização")
	default:
		fmt.Println("Outro dia Com inicialização")
	}

}

func main() {
	normal()
	comInicializacao()
	comBooleano()
	switchCase()
	switchCaseComVariavel()
	switchCaseComInicializacao()
}
