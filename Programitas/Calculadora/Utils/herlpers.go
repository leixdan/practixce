package utils

import (
	"fmt"
	"os"
)

//ImprimirResultado imprime en pantalla el resultado de la operación
func ImprimirResultado (a, b, resultado float64){
	fmt.Printf("El resultado es: %f\n", resultado)
}

//TerminarPrograma finaliza el programa al encontrar un error
func TerminarPrograma(){
	fmt.Println("Se termina el programa")
	os.Exit(0)
}