package utils

import (
	"fmt"
	"os"
)

//ImprimirResultado imprime en pantalla el resultado de la operación
func ImprimirResultado (a, b, resultado int){
	fmt.Printf("El resultado es: %d\n", resultado)
}

func TerminarPrograma(){
	fmt.Println("Se termina el programa")
	os.Exit(0)
}