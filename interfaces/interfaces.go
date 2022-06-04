package main

import (
	"fmt"
)

func contar_caracteres(a any) int {
	texto := a.(string)
	var soma_dos_caracters int = 0
	for _, _ = range texto {
		soma_dos_caracters += 1
	}

	return soma_dos_caracters
}

func somar(a any) int {
	lista := a.([]int)

	var soma int
	for _, item := range lista {
		soma += item
	}

	return soma
}

func main() {
	//resultado := somar([]int{1, 2, 3, 4})
	total_caracteres := contar_caracteres("abc def")
	fmt.Println(total_caracteres)

	total_caracteres = contar_caracteres(123)
	fmt.Println(total_caracteres)
}
