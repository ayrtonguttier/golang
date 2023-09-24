package functions

type pessoa struct {
	Nome string
}

func NewPessoa(nome string) pessoa {
	return pessoa{Nome: nome}
}
