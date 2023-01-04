package main

import "fmt"

func main() {

	// Map é uma estrutura guarda valores, sendo por chave -> valor
	// Chaves e valores devem ser todos do mesmo tipo, ou seja, se você usou string no tipo de chave, o seu valor vai ter que ser string também
	// Dentro de [] é o tipo da chave e fora tipo do dado

	user := map[string]string{
		"nome":      "Luiz",
		"sobrenome": "Silva",
	}

	fmt.Println(user) // map[nome:Luiz sobrenome:Silva]

	// Para acessar um valor dentro do map
	fmt.Println(user["nome"]) // Luiz

}
