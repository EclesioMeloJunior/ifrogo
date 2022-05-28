package desafio

import (
	"fmt"
	"testing"
)

func TestDesafio(t *testing.T) {
	lista := []string{}
	fmt.Println(lista)
	nos := CriarNos(lista)

	primeiro := nos[0]

	primeiro.Apontar(nos[1:]...)

	esperado := "fernando"
	resultado := Imprimir(primeiro)

	if resultado != esperado {
		t.Errorf("esperado %s, o que recebi: %s",
			esperado, resultado)
	}
}
