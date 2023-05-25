package main

import "fmt"

func main() {
	const month = 12
	months := []string{
		"",
		"Janeiro",
		"Fevereiro",
		"Março",
		"Abril",
		"Maio",
		"Junho",
		"Julho",
		"Agosto",
		"Setembro",
		"Outubro",
		"Novembro",
		"Dezembro",
	}
	fmt.Println(months[month])
}