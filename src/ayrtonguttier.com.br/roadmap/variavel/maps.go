package variavel

import (
	"fmt"

	"ayrtonguttier.com.br/roadmap/cli"
)

func RunMaps() {
	cli.Print("Maps")
	emptyMap()
	mapSemMake()
	deleteMapItem()
	consultaElemento()
}

func emptyMap() {
	//map[Key]Value
	//map é inicializado pela função make
	m := make(map[string]int)

	m["ayrton"] = 29
	fmt.Println(m)
}
func mapSemMake() {
	m := map[string]int{
		"ayrton": 29,
	}

	m["guttier"] = 30
	fmt.Println(m)
}

func deleteMapItem() {
	m := newMap()
	delete(m, "guttier")
	fmt.Println(m)

}

func consultaElemento() {
	m := newMap()
	printElemento(m, "ayrton")
	printElemento(m, "zé")


}

func printElemento(m map[string]int, key string) {

	elem, ok := m[key]
	if ok {
		fmt.Printf("Elemento encontrado(%s): %d\n", key, elem)
	}else{
		fmt.Printf("Elemento não encontrado(%s)\n", key)
	}

}

func newMap() map[string]int {
	return map[string]int{
		"ayrton":  29,
		"guttier": 30,
	}
}
