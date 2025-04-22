//Ejercicios mamalones de punteros porque estoy bien wey y se me olvidan las funciones tambien

package main

import 
	"fmt"


//Ej 1
func duplicar(n int) int { //Ej 1
	return n * 2
}
//Ej 2
func triplicarPuntero(n *int) { //Ej 2
	*n = *n * 3
}
//Ej 3
func cambiarNombre(name *string, nuevo string){
	*name = nuevo
}
//Ej 4
func intercambiar(a, b *int){
	*a, *b = *b, *a
}
//Ej 5
type Persona struct{
	nombre string
	edad int
}
func (p *Persona) cumpleaños (){
	p.edad++
}



func main() {

	//Variables

	var dup int
	var name string = "Dolores"
	var pNombre string

	//Ejercicio 1
	fmt.Println("Ingresa el número que deseas duplicar: ")
	fmt.Scan(&dup)
	resultado := duplicar(dup)
	fmt.Println("El resultado es:", resultado)
	fmt.Println("=================================================")

	//Ejercicio 2
	fmt.Println("Número triplicado por puntero: ")
	triplicarPuntero(&dup)
	fmt.Println("El resultado es:", dup)
	fmt.Println("=================================================")

	//Ejercicio 3 - Nuevo nombre con entrada del usuario
	fmt.Printf("Por favor actualiza el nombre: %s\n",name)
	fmt.Println("Ingresa el nuevo nombre a actualizar en el espacio en memoria: ")
	fmt.Scan(&pNombre)
	cambiarNombre(&name, pNombre)
	fmt.Println("Cool, ahora el nombre en la variable name es:", name)
	fmt.Println("=================================================")

	//Ejercicio 4 - Intercambiar dos valores (swap con punteros)
	a:= 10
	b:= 20

	fmt.Printf("Tenemos la variable a con valor: %d y la variable b con valor: %d\n", a, b)
	fmt.Println("Haz que se intercambien neita")

	intercambiar(&a, &b)

	fmt.Println("El nuevo valor de a: ",a)
	fmt.Println("El nuevo valor de b: ",b)

	fmt.Println("=================================================")

	//Ejercicio 5 - Incrementar edad en una estructura

	p := Persona{
		nombre: "Florecita",
		edad: 23,
	}
	
	fmt.Println(p)
	fmt.Println("Vamos a actualizar a Florecita, el ya cumple años ñerito")
	p.cumpleaños()
	fmt.Println(p)

}
