package auxiliar

import "fmt"

// Caso o nome da func seja em maiúscula, quer dizer que essa função pode ser exportada
// Caso o nome da func seja em minúscula, quer dizer que essa função só pode ser chamada em seu próprio arquivo ou no mesmo pacote
// Criando um comentário antes da linha que cria a função, isso serve como um docs

// Func que inicia nosso auxiliar "fake"
func Iniciar() {
	fmt.Println("Iniciando Auxiliar...")
	fmt.Println("Auxiliar Iniciado Com Sucesso!")
}
