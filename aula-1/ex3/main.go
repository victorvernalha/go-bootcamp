package main

import "fmt"

func main() {
	var nome string
	var sobrenome string
	var idade int
	idade = 6
	var licenca_para_dirigir = true
	var estatura_da_pessoa int
	quantidadeDeFilhos := 2

	// Print values to avoid unused varuabkle warnings
	fmt.Println(nome, sobrenome, idade, licenca_para_dirigir, estatura_da_pessoa, quantidadeDeFilhos)
}
