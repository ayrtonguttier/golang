package functions

import (
	"fmt"

	"ayrtonguttier.com.br/roadmap/cli"
)

type FuncaoSoma func(int, int) int

func RunFunctions() {
	cli.Print("Functions")
	funcaoSimples()
	funcaoComParametros(1, 1)
	retornoPrimeiraFuncao := funcaoComRetorno(2, 2)
	fmt.Println(retornoPrimeiraFuncao)
	retornoSegundaFuncao := funcaoComRetornoNomeado(2, 2)
	fmt.Println(retornoSegundaFuncao)
	soma, potencia := multiplosRetornos(3, 3)
	fmt.Println("Resultado multiplos retornos soma: ", soma, " potencia: ", potencia)

	pSoma := &soma

	fmt.Println("Ponteiro da variavel soma", pSoma)
	funcaoComPonteiro(pSoma)
	fmt.Println("Novo valor de soma: ", soma)

	var teste = func(a, b int) int {
		return a + b
	}

	fmt.Println("Função anonima com resultado: ", teste(2, 2))

	func(a int) {
		fmt.Println("Outra forma de função anonima ", a)
	}(20)

	func() {
		fmt.Println("função clojure acessa membros da função pai ", soma)
	}()

	funcaoSoma := FuncaoSoma(func(x, z int) int {
		return x + z
	})

	fmt.Println("Soma com typed function ", funcaoSoma(2, 2))	
}

func funcaoSimples() {
	fmt.Println("Faz qualquer coisa")
}

func funcaoComParametros(a, b int) {
	fmt.Println(a + b)
}

func funcaoComRetorno(a, b int) string {
	return fmt.Sprintf("Função com retorno a: %d b: %d soma: %d", a, b, a+b)
}

func funcaoComRetornoNomeado(a, b int) (resultado string) {
	resultado = fmt.Sprint("Funcão com retorno nomeado: ", a*b)
	return
}

func multiplosRetornos(a, b int) (multiplicacao, potencia int) {
	multiplicacao = a * b

	potencia = 1
	for i := 0; i < b; i++ {
		potencia = potencia * a
	}
	return
}

func funcaoComPonteiro(a *int) {
	soma, _ := multiplosRetornos(2, 2)
	*a = soma
}
