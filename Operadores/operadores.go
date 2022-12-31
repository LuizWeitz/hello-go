package main

import "fmt"

func main() {

	soma := 1 + 2
	subtracao := 1 - 9
	divisao := 9 / 2
	multiplicacao := 8 * 8
	restoDeDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDeDivisao)

	// Só podemos realizar operações com os mesmo tipo de dado e de bytes

	var n1 int16 = 10
	var n2 int16 = 12

	var soma2 int16 = n2 + n1

	fmt.Println(soma2)

	// Atribuição de valores

	// =, usado para atribuir valor quando informamos o tipo
	var s1 string = "Hello"

	// :=, usado para atribuir valor quando não informamos o seu tipo

	s2 := "World"

	fmt.Println(s1 + " " + s2)

	// Operadores relacionais
	// Sempre retornam um valor boolean

	fmt.Println(1 >= 2) // false
	fmt.Println(1 == 2) // false
	fmt.Println(1 <= 2) // true
	fmt.Println(1 > 2)  // false
	fmt.Println(1 < 2)  // true
	fmt.Println(1 != 2) // true

	// Operadores lógicos

	// && -> e, retorna true se todas as condições forem verdadeiras
	fmt.Println(1 > 2 && 2 < 1) // false

	// || -> ou, retorna true se alguma condição for verdadeira
	fmt.Println(1 == 2 || 1 < 2) // true

	// Operadores uniários

	n3 := 10

	// Add +1 ao valor
	n3++ // 11

	// Pega o valor atual e add algum valor
	n3 += 10 // 20

	// Pega o valor atual e subtrai
	n3 -= 10 // 0

}
