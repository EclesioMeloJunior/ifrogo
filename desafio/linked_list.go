package desafio

import "fmt"

func CriarNos(letras []string) []*No {
	nos := []*No{}

	for _, letra := range letras {
		no := &No{Valor: letra}
		nos = append(nos, no)
	}

	return nos
}

type No struct {
	Valor string
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

func Imprimir(n *No) string {
	completo := ""

	for n != nil {
		fmt.Printf("%s -> ", n.Valor)
		completo += n.Valor
		n = n.Braco
	}
	fmt.Println()
	return completo
}
