package main

import (
	"codif/utils"
	"fmt"
)

func main() {

	fmt.Println("Bienvenido al codificador Gaskull 1.0")
	fmt.Println("Ingresa la palabra que quieres modificar")
	var entrada string
	fmt.Scan(&entrada)
	//Transformamos el string a rune
	SEntrada := []rune(entrada)

	utils.ModificarEntrada(&SEntrada)

}
