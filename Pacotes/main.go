package main

// Importando o que vamos usar
import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Bem-Vindo Aos Pacotes Com Go...")

	auxiliar.Iniciar()

	// Usando um pacote externo de validação de endereço email
	validacaoEmail := checkmail.ValidateFormat("teste@gmail.com")

	// Vai retornar nil caso o endereço email seja válido
	fmt.Println(validacaoEmail)
}
