package desafio

import (
	"testing"
)

func TestLinkedList(t *testing.T) {

	a := &No{Valor: ""}
	b := &No{Valor: ""}
	c := &No{Valor: ""}
	d := &No{Valor: ""}

	a.Apontar(b, c, d)
	resultado := Imprimir(a)
	esperado := "eclesio"

	if resultado != esperado {
		t.Fail()
	}
}
