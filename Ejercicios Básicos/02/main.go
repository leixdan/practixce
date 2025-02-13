//Generar el inventrio de un estudio de tatuajes, se debe poder agregar y eliminar el material disponible, y mostrar el inventario actual.
/*
1. Agregar mateial según el tipo de material
2. Mostrar el inventario completo
3. Buscar el material por nombre y mostrar la cantidad disponible e información
4. Modificar existencias
5. Eliminar material
*/

package inventory

import (
	"fmt"
)

// Función para agregar material al inventario
func addmaterial(material string, quantity int) { //Se recibe por el usuario el material y la cantidad
	inventory[material] = quantity //Se agrega el material al inventario con la cantidad, si ya existe se actualiza la cantidad
}

// Función para mostrar el inventario completo
func showinventory() {
	for material, quantity := range inventory { //Se recorre el inventario y se obtiene el material y la cantidad
		fmt.Printf("En la bodega tenemos: %d - %s", quantity, material) //Se imprime el material y la cantidad
	}
}

// Función para buscar material por nombre
func searchmaterial(material string) {
	if _, ok := inventory[material]; ok { //Se busca el material en el inventario
		fmt.Printf("El material %s tiene una cantidad de %d", material, inventory[material]) //Se imprime el material y la cantidad
	} else {
		fmt.Println("El material no se encuentra en el inventario") //Si no se encuentra el material se imprime un mensaje
	}
}

// Función para modificar existencias
func modifyquantity(material string, quantity int) {
	inventory[material] = quantity //Se modifica la cantidad del material
}

// Función para eliminar material
func deletematerial(material string) {
	delete(inventory, material) //Se elimina el material del inventario
}

func main() {
	inventory := make(map[string]int) //Se crea el inventario como un mapa de string e int

	fmt.Println("Bienvenido al inventario del estudio")
	fmt.Println("1. Agregar material")
	fmt.Println("2. Mostrar inventario")
	fmt.Println("3. Buscar material")
	fmt.Println("4. Modificar existencias")
	fmt.Println("5. Eliminar material")

	var option int      //Se declara la variable para la opción del usuario
	fmt.Scanln(&option) //Se recibe la opción del usuario y se guarda en la variable "option"

	switch option { //Se crea un switch para las opciones del usuario
	case 1: //En caso de que la opción sea 1
		var material string                //Se declara la variable para el material
		var quantity int                   //Se declara la variable para la cantidad
		fmt.Println("Ingrese el material") //Se pide al usuario que ingrese el material
	}
}
