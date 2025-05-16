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

	//Lo mandamos a codificar:
	utils.Codificar(SEntrada,3)
	//Mandamos el slice y recibimos un string de:
	utils.ConvertirEntrada(SEntrada)
	//Mandado a decodificar
	utils.Decodificar(SEntrada,3)
	//Mandamos el slice y recibimos un string de:
	utils.ConvertirEntrada(SEntrada)
}
