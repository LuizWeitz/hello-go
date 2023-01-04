package main

import "fmt"

func diaDaSemana(numero int) string {

	switch numero {
	case 0:
		return "Domingo"
	case 1:
		return "Segunda-Feira"
	case 2:
		return "Terça-Feira"
	case 3:
		return "Quarta-Feira"
	case 4:
		return "Quinta-Feira"
	case 5:
		return "Sexta-Feira"
	case 6:
		return "Sábado"
	default:
		return "Número Inválido"
	}

}

// Podemos também fazer da seguinte forma
// Dessa forma atribuímos o valor do return direto em uma variável e retornamos essa variável no final

func diaDaSemana2(numero int) string {

	var diaDaSemana string

	switch numero {
	case 0:
		diaDaSemana = "Domingo"
	case 1:
		diaDaSemana = "Segunda-Feira"
	case 2:
		diaDaSemana = "Terça-Feira"
	case 3:
		diaDaSemana = "Quarta-Feira"
	case 4:
		diaDaSemana = "Quinta-Feira"
	case 5:
		diaDaSemana = "Sexta-Feira"
	case 6:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número Inválido"
	}

	return diaDaSemana

}

func main() {

	fmt.Println(diaDaSemana(1))
	fmt.Println(diaDaSemana(4))
	fmt.Println(diaDaSemana(10))

	fmt.Println(diaDaSemana2(13))
	fmt.Println(diaDaSemana2(1))
	fmt.Println(diaDaSemana2(0))

}
