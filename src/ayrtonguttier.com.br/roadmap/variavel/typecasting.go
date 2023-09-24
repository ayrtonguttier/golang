package variavel

import (
	"fmt"
	"strconv"

	"ayrtonguttier.com.br/roadmap/cli"
)

func RunTypeCasting() {
	cli.Print("Type Casting")
	integerToFloat()
	floatParaInteiro()
	stringAndBytes()
	stringParaInteiro()
	floatAndIntegerDivision()
}

func integerToFloat() {
	i := 10
	f := float64(i)

	printValores(i, f)
}

func floatParaInteiro() {
	f := 10.15
	i := int(f) // Perde a precisão

	printValores(i, f)
}

func stringAndBytes() {
	n := "Ayrton"
	b := []byte(n)
	n2 := string(b)

	fmt.Println(n)
	fmt.Println(b)
	fmt.Println(n2)
}

func stringParaInteiro() {
	valor := "111"
	i, _ := strconv.ParseInt(valor, 10, 0)
	iBinary, _ := strconv.ParseInt(valor, 2, 0)
	fmt.Println("Base 10: ", i, " Base 2: ", iBinary)
}

func floatAndIntegerDivision() {
	i := 6 / 3   // a é inteiro
	f := 8.2 / 2 // f é float

	// No caso de variável da erro e precisa de conversão
	printValores(i, f)
}

func printValores(i int, f float64) {
	fmt.Printf("Valor inteiro: %d Valor float: %f\n", i, f)
}
