package operacao


import (
	"fmt"
	"projeto/ctt"
)



func AdicionaContato(c ctt.Contato, a *[5]ctt.Contato) {
	for ind, contato := range a {
		if (contato == ctt.Contato{}) {
			a[ind] = c
			break
		}
	}


}

func ExcluiContato(a *[5]ctt.Contato) {
	if (a[0] == ctt.Contato{}) {
		fmt.Println("Lista de contatos vazia!")
	}

	for ind, contato := range a {
		if (contato == ctt.Contato{}) {
			a[ind-1] = ctt.Contato{}
		}
	}
}

func EditaEmail(indice int, novoEmail string , a *[5]ctt.Contato) {
	if indice < 0 && indice >=len(a) {

	}
	if a[indice] == (ctt.Contato[]) {

	}

	a[indice].Email = novoEmail
}