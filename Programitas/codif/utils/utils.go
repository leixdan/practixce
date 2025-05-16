package utils

import (
	"fmt"
)

//Codificar obtiene el slice y lo modifica
func Codificar(s []rune, shift int){

	for i := range s {
		s[i] = s[i] + rune(shift)
	}

}

// ModificarEntrada obtiene la respuesta del usuario y la modifica
func ConvertirEntrada(r []rune) {
	fmt.Println(string(r))
}

//Decodificar obtiene el slice y lo modifica
func Decodificar(s []rune, shift int){

	for i := range s {
		s[i] = s[i] - rune(shift)
	}

}