package main

import "fmt"

func main() {

	idade := 18

	if idade >= 18 {
		fmt.Println("Pessoa Maior de Idade")
	} else {
		fmt.Println("Pessoa Menor de Idade")
	}

	// Podemos iniciar uma variável direto no if/else
	// Iniciando uma variável e já aplicando uma condição para ela
	// Não conseguimos usar essa variável fora do if/else que ela pertence
	if idade2 := 8; idade2 >= 18 {
		fmt.Println("Pessoa Maior de Idade")
	} else {
		fmt.Println("Pessoa Menor de Idade")
	}

}
