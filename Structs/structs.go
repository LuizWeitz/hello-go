package main

import "fmt"

// Nesse caso estou usando uma união de structs

type usuario struct {
	nome     string
	email    string
	idade    int8
	endereco endereco
}

type endereco struct {
	rua    string
	numero int8
}

func main() {

	// Criando uma variável com modelo do struct criado
	var user usuario

	// Atribuindo valores ao usuário instanciado
	user.nome = "Luiz"
	user.email = "luiz@email.com"
	user.idade = 21

	// Atribuindo valores ao struct endereço que tem dentro do struct do usuário
	user.endereco.rua = "Rua das Palmeiras"
	user.endereco.numero = 23

	fmt.Println(user)

	// Também podemos atribuir esses valores direito na chamada da variável do tipo do usuário
	endereco := endereco{"Rua dos Palmitos", 123}

	user2 := usuario{"Luiz", "luiz@email2.com", 18, endereco}

	fmt.Println(user2)

}
