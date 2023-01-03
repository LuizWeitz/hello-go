package main

import "fmt"

func main() {

	// Trabalhando com Arrays

	// Basicamente a estrutura é composta por nome, quantidade de itens e o tipo desses itens
	// Não podemos ter tipos diferentes em um array, somente um tipo de dados
	var array1 [10]int

	fmt.Println(array1) // [0 0 0 0 0 0 0 0 0 0]

	// Para acessar uma posição e colocar valor nela
	// Lembrando array começa sempre na posição 0
	array1[0] = 1
	array1[1] = 2

	fmt.Println(array1) // [1 2 0 0 0 0 0 0 0 0]

	// Iniciando um array com valores
	// Qunado iniciamos com valores, não precisamos informar a palavra var
	array2 := [3]string{"P0", "P1", "P2"}

	fmt.Println(array2) // [P0 P1 P2]

	// Caso queremos pegar o tamanho de um array
	fmt.Println(len(array2)) // 3

	// Trabalhando com Slices
	// Slices são parecidos com arrays, só que mais flexiveis para trabalhar

	// Basicamente a estrutura é composta por nome, [] e o seu tipo
	// Slice não precisamos informar uma quantidade
	slice := []int{}

	fmt.Println(slice) // []

	// Para colocamos um valor em uma posição
	// Função append, pega o slice que existe, add o valor que passamos e retorna um slice novo
	// Como nesse caso estamos se referindo a um slice que já existe, já vai substituir ele
	slice = append(slice, 1)

	fmt.Println(slice) // [1]

	// Caso queremos inicar um slice com valores
	slice2 := []int{123, 124, 125, 126, 127}

	fmt.Println(slice2) // [123 124 125 126 127]

	// Caso queremos pegar os valores de uma faixa de posições
	// Nesse caso ele vai pegar os valores da possição 1 a 2, mesmo passando de 1:3, parando antes o index 3
	fmt.Println(slice2[1:3]) // [124 125]

	// Caso queremos criar um slice com os valores de um array existente
	slice3 := array2[1:2]

	fmt.Println(slice3) // [P1]

	// Caso queremos pegar o tamanho de um slice

	fmt.Println(len(slice3)) // 1

}
