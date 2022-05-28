package linkedin

import (
	"testing"
)

func TestLinkedList(t *testing.T) {

	a := &No{Valor: 1}
	b := &No{Valor: 2}
	c := &No{Valor: 3}
	d := &No{Valor: 4}

	a.Apontar(b, c, d)
	Imprimir(a)
}
