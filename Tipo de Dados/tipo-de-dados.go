package main

import (
	"errors"
	"fmt"
)

func main() {

	// Atenção : nesses exemplos estou iniciando os tipos com algum valor, mas não é obrigatório

	// Tipo int
	// Seguindo a ordem, primeiro o tipo int e depois a quantidade de bytes
	// Sendo int8, int16, int32 e int64
	// Caso não informado a quantidade de bytes, automáticamente será usado o tipo de arquitetura da máquina atual

	// Alguns exemplos de int
	var n1 int64 = 1000000
	var n2 int16 = 1000

	fmt.Println(n1, n2)

	// Podemos também atribuir valores negativos
	var n3 int16 = -3403
	var n4 int8 = -14

	fmt.Println(n3, n4)

	// Também temos o tipo uint, que é um int sem sinal, só podemos usar valores sem sinal, basicamente só valores positivos
	var n5 uint16 = 123

	fmt.Println(n5)

	// Tipo float, segue mesma lógica que o tipo int, só que aceita ponto, valores quebrados
	// Só que aceita apenas tipos 16 e 32 bytes
	var n6 float32 = 12.98

	fmt.Println(n6)

	// Tipo string
	var s1 string = "Hello"

	fmt.Println(s1)

	// Tipo boolean
	// Caso não colocado nenhum valor, por padrão é false
	var b1 bool = true

	fmt.Println(b1)

	// Em Go temos o tipo error
	// Por padrão retornar nil
	// Caso queremos atribuir um valor, temos que importar o pacote errors
	var e1 error = errors.New("Um Erro")

	fmt.Println(e1)

}
