package main

import (
	"fmt"
	"sync"
)

// Definición de una estructura (struct)
type Persona struct {
	Nombre string
	Edad   int
}

func main() {
	// Declaración de variables de distintos tipos
	var x int = 10
	var nombre string = "Juan"
	var pi float64 = 3.14
	var activo bool = true

	// Arreglo (array) de enteros
	var arreglo [3]int = [3]int{1, 2, 3}

	// Slice de enteros
	slice := []int{4, 5, 6}

	// Mapa (map) de string a int
	mapa := map[string]int{"uno": 1, "dos": 2}

	// Crear un mapa vacío
	mimapa := make(map[string]int)

	// Agregar elementos al mapa
	mimapa["r1"] = 10
	mimapa["r2"] = 20

	// Struct
	juan := Persona{
		Nombre: "Juan",
		Edad:   30,
	}

	// Go routine
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Esta es una Go routine")
	}()

	// Llamada a función
	resultado := sumar(5, 7)

	// Llamadas a funciones con iteración sobre slices y mapas
	slice = agregarElementoSlice(slice, 7)
	fmt.Println("Slice después de agregar un elemento:", slice)

	// Iterar sobre slice
	iterateOverSlice(slice)

	// Iterar sobre mapa
	iterateOverMap(mapa)

	// Impresión de todos los elementos
	fmt.Println("Variable x:", x)
	fmt.Println("Nombre:", nombre)
	fmt.Println("Valor de pi:", pi)
	fmt.Println("Activo:", activo)

	// Arreglo
	fmt.Println("Arreglo:", arreglo)

	// Mapa
	fmt.Println("Mapa:", mapa)
	fmt.Println("Contenido del mapa:", mimapa)

	// Struct
	fmt.Println("Persona:", juan)

	// Resultado de la función
	fmt.Println("Resultado de la suma:", resultado)

	// Esperar a que finalice la Go routine
	wg.Wait()
}

// Función simple que suma dos números
func sumar(a, b int) int {
	return a + b
}

// Función que agrega un elemento a un slice
func agregarElementoSlice(slice []int, valor int) []int {
	slice = append(slice, valor) // Se agrega el valor al slice
	return slice
}

// Función que recorre un slice e imprime cada uno de sus elementos
func iterateOverSlice(slice []int) {
	fmt.Println("Iterando sobre el slice:")
	for _, v := range slice {
		fmt.Println(v)
	}
}

// Función que recorre un mapa e imprime clave y valor
func iterateOverMap(mapa map[string]int) {
	fmt.Println("Iterando sobre el mapa:")
	for clave, valor := range mapa {
		fmt.Printf("Clave: %s, Valor: %d\n", clave, valor)
	}
}
