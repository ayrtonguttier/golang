# Anotações sobre functions

## Exports

Exports funcionam com base no primeiro caracter do nome do que está sendo exportado.

Algumas coisas que podem ser exportadas são:
- functions
- types
- structs
- variaveis do struct

Importante ressaltar, mesmo que um desses itens não sejam importados, eles ainda podem ser acessados através de um outro item que seja exportado.


``` golang

type pessoa struct{
    Nome string
    sobrenome string
}

func NewPessoa(nome, sNome string){
    return pessoa{ Nome:nome, sobrenome: sNome }
}

```

Nesse caso a struct pessoa não pode ser iniciada em outro package sem a utilização da função `NewPessoa`.

A partir do momento em que a função `NewPessoa` é utilizada um novo valor de pessoa é criado e retornado.

``` golang

func main(){
    p := NewPessoa()
    fmt.Println(p.Nome)
}
```
Outro ponto importante. Variáveis dentro de structs também seguem a mesma regra de export. Nesse caso a variável `Nome` pode ser acessada mas a `sobrenome` não.
