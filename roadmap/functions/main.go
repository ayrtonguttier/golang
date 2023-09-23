package main

import "fmt"

func main() {
	funcaoSimples()
	funcaoComParametros(1,1)
	fmt.Println(funcaoComRetorno(2,2))
}

func funcaoSimples() {
	fmt.Println("Faz qualquer coisa")
}

func funcaoComParametros(a, b int) {
	fmt.Println(a + b)
}


func funcaoComRetorno(a,b int) string{
	return fmt.Sprintf("Função com retorno a: %d b: %d soma: %d", a, b, a+b)
}