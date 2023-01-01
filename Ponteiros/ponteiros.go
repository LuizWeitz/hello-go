package main

import "fmt"

// Ponteiro são como variáveis que salva o endereço de memória de alguma coisa

func main() {

	var v1 int
	var p1 *int // Guarda o valor de memória de um valor inteiro

	v1 = 100
	p1 = &v1 // Atribuindo o caminho de memória da variável v1 para o ponteiro

	fmt.Println(v1, p1)

	// Caso queremos ler o valor que tem no caminho apontado pelo ponteiro
	fmt.Println(*p1)

	// Por que usar ponteiro?
	// Vamos supor que temos duas variáveis

	var v2 int = 20
	var v3 int = v2

	// Caso alteremos o valor de v2, o valor de v3 vai continuar 20, pois apenas fizemos uma copia de v2 para v3 na criação de v3
	v2++

	fmt.Println(v2, v3)

	// Nesse caso que entra o uso de ponteiro, com ele podemos alterar o valor de v2 que v3 muda
	// Pois v3 vai ser o caminho de memória de v2, nesse caso só precisamos acessar o valor desse caminho
	// Assim vamos obter o novo valor de v2 para v3

	// Bora refatorar, para ser mais didático, vou criar novas variáveis

	var v4 int = 10
	var v5 = &v4

	fmt.Println(v4, v5)

	// Alterando o valor de v4
	v4++

	// Caso acessemos o valor de v5 via caminho de memória de v4, seu valor vai ser o novo de v4
	fmt.Println(v4, *v5)

}
