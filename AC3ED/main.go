package main

import (
	"fmt"
	"bufio"
	"os"
	"projeto/ctt"
	"projeto/operacao"

)


func main() {
	var contatos [5]ctt.Contato
	var nome, email, opcao string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			operacao.AdicionaContato(ctt.Contato{Nome: nome, Email: email}, &contatos)
		case "2":
			operacao.ExcluiContato(&contatos)
		default:
			return
		}

		fmt.Println(contatos)
	}

}
