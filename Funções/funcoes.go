package main

import "fmt"

// Iniciamos uma função usando func
// Entre parênteses passamos os parâmetros e logo após o tipo que nossa função vai retornar

// Função que recebe dois valores tipo int64 e retornar a soma deles
func somar(n1 int64, n2 int64) int64 {
	return n1 + n2
}

// Uma função pode ter vários retornos, só precisamos informar os tipos dos retornos depois do parênteses
// Função que recebe dois int64 e retorna sua soma, subtração, multiplicação e divisão
func operacoesMatematicas(n1 int64, n2 int64) (int64, int64, int64, float64) {

	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2

	return soma, subtracao, multiplicacao, float64(divisao)

}

func main() {

	// Declarando uma variável para nossa função
	// Chamando a func criada, passando os valores para os parâmetros
	resultado := somar(1, 2)

	// Declarando cada variável a cada retorno da função
	// Caso queremos ignorar algum retorno, colocamos apenas _
	resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao := operacoesMatematicas(8, 2)

	fmt.Println(resultado)
	fmt.Println(resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao)
}
