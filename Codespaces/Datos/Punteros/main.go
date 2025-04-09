package main

import "fmt"

func updateValue (p *int, newValue int){
	*p = newValue
}

func main (){

//Puntero básico

var ogName string = "Karla"
var newName *string = &ogName

fmt.Println("Su nombre es: ", ogName)
fmt.Println("Se actualizará el espacio de memoria: ",newName)
fmt.Println("Valor de nuevo nombre sin actualizar: ", *newName)

*newName = "Florecita"
fmt.Println("El nombre actualizado es: ",ogName)


//Puntero desde función
	var a int = 123
	var b *int = &a //Para una variable-puntero, siempre se declara con * en el tipo de dato

	fmt.Println("El valor de <a> original es: ",a)
	fmt.Println("El espacio en memoria asignado a <b> es: ",b)
	
var newValue int
fmt.Println("Ingresa el nuevo valor a modificar chiko: ")
fmt.Scan(&newValue)

updateValue(b, newValue)

fmt.Println("El nuevo valor de <a> modificado es: ", a)
fmt.Println("El nuevo valor de <a> desde <*b> es: ", *b)

}