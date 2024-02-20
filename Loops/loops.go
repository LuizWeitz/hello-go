package main

import (
	"fmt"
	"time"
)

func main() {

	i := 0

	// Enquanto a condição for menor que 10, executa o bloco de código
	// Nesse caso, vai iniciar no 1, assim imprimindo até 10
	for i < 10 {

		time.Sleep(time.Second) // "Dorme por um segundo"
		fmt.Println("Adicionado +1 a i")
		i++ // Adiciona +1 a i

		fmt.Println(i)

	}

	// Estrutura do for mais simples e direta
	// Nesse caso, vai até o 4, imprimindo até o 4
	for j := 0; j < 5; j++ {
		fmt.Println("Adicionado +1 a j")
		time.Sleep(time.Second)
		fmt.Println(j)

	}

	// Podemos usar tanto Arrays ou Slices
	nomes := [3]string{"Gabriel", "Thanos", "Steve"}

	// Primeiro parâmetro é o indice e o segundo o valor
	for indice, nome := range nomes {
		fmt.Println(indice, nome) // Retornar posição e o valor
	}

	// Caso eu queria apenas retornar o valor do indice, mas sem o indice
	for _, nome := range nomes {
		fmt.Println(nome)
	}

}
