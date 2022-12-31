package main

import "fmt"

func main() {

	// Para criar uma variável, podemos atribuir seu tipo ou não

	// Atribuindo seu tipo
	var v1 string = "v1"
	fmt.Println(v1)

	// Sem atribuir seu tipo
	v2 := "v2"
	fmt.Println(v2)

	// Também podemos criar várias variáveis de uma vez só, com ou sem seus tipos

	// Com seus tipos
	var (
		v3 string = "v3"
		v4 string = "v4"
	)

	fmt.Println(v3, v4)

	// Sem seus tipos
	v5, v6 := "v5", "v6"

	fmt.Println(v5, v6)

	// Podemos usar constantes também

	const c1 string = "c1"
	fmt.Println(c1)

}
