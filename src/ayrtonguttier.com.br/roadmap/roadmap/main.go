package main

import (
	"fmt"

	"ayrtonguttier.com.br/roadmap/conditionals"
	"ayrtonguttier.com.br/roadmap/functions"
	"ayrtonguttier.com.br/roadmap/variavel"
)

func main() {
	conditionals.RunConditionals()
	functions.RunFunctions()
	functions.RunFilter()
	p := functions.NewPessoa("Ayrton")
	fmt.Println(p.Nome)
	variavel.RunVariable()
	variavel.RunTypeCasting()
	variavel.RunMaps()
}
