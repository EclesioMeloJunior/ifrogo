package linkedin

import "fmt"

type No struct {
	Valor int
	Braco *No
}

func (n *No) Apontar(outros ...*No) {
	if len(outros) == 0 {
		return
	}

	primeiro := outros[0]
	n.Braco = primeiro
	primeiro.Apontar(outros[1:]...)
}

func Imprimir(n *No) {
	for n != nil {
		fmt.Printf("%d ->", n.Valor)

		n = n.Braco
	}

	fmt.Println()
}
