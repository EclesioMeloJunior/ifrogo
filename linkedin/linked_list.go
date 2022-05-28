package linkedin

import "fmt"

type No struct {
	Valor int
	Braco *No
}

func (n *No) Apontar(outro *No) {
	n.Braco = outro
}

func Imprimir(n *No) {
	for n != nil {
		fmt.Printf("%d -> ", n.Valor)

		n = n.Braco
	}

	fmt.Println()
}
