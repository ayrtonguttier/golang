package cli

import "fmt"

func Print(txt string) {
	fmt.Println(completarTracos(txt))
}

func completarTracos(txt string) string {
	total := 100
	lTxt := len(txt)
	qtdTracos := (total - lTxt) / 2

	var tracos string

	for i := 0; i < qtdTracos; i++ {
		tracos += "-"
	}

	if lTxt % 2 != 0{
		txt += " "
	}


	return tracos + " " + txt + " " + tracos
}
