package main

import "fmt"

/*________________________*/
/* Tipos de Datos Básicos */

func main() {
	// Tipos primitivos
	var entero int = 42
	var flotante float64 = 3.14
	var cadena string = "Hola, Go!"
	var booleano bool = true

	fmt.Println("Entero:", entero)
	fmt.Println("Flotante:", flotante)
	fmt.Println("Cadena:", cadena)
	fmt.Println("Booleano:", booleano)
}

/*________________________*/
/* Arreglos */

func main() {
	// Arreglo con tamaño fijo
	var arreglo [3]int
	arreglo[0] = 10
	arreglo[1] = 20
	arreglo[2] = 30

	fmt.Println("Arreglo:", arreglo)
}

/*________________________*/
/* Slices */

func main() {
	// Creación y manipulación de slices
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice1:", slice1)

	// Añadir un elemento al slice
	slice1 = append(slice1, 6)
	fmt.Println("Slice1 después de append:", slice1)

	// Subslice (subconjunto del slice original)
	slice2 := slice1[2:5]
	fmt.Println("Slice2 (subslice):", slice2)
}

/*________________________*/
/* Mapas */

// Mapa básico
func main() {
	// Mapa de clave-valor
	mapa := map[string]int{
		"manzana": 5,
		"banana":  8,
		"cereza":  3,
	}

	fmt.Println("Mapa:", mapa)

	// Accediendo a un valor por clave
	cantidadManzanas := mapa["manzana"]
	fmt.Println("Cantidad de manzanas:", cantidadManzanas)
}

// Mapa con Slice como clave
func main() {
	// Mapa con slice como clave (notar que Go no permite slices como clave directa, entonces usamos array fijo)
	mapa := make(map[[2]int]string)

	// Definir clave de array
	clave := [2]int{1, 2}
	mapa[clave] = "Valor asociado a [1, 2]"

	fmt.Println("Mapa con clave de array:", mapa)
}

// Mapa con función
func main() {
	// Mapa donde el valor es una función
	operaciones := map[string]func(int, int) int{
		"sumar":  func(a, b int) int { return a + b },
		"restar": func(a, b int) int { return a - b },
	}

	// Llamada a la función dentro del mapa
	resultado := operaciones["sumar"](5, 3)
	fmt.Println("Resultado de sumar:", resultado)
}

// Mapa con estructura
func main() {
	// Mapa con estructura como valor
	mapa := map[string]Persona{
		"juan": Persona{
			Nombre: "Juan",
			Edad:   30,
			Ciudad: "Madrid",
		},
	}

	fmt.Println("Mapa con estructura:", mapa)
}


// Mapa desde cero
func main() {
	// Crear un mapa vacío
	mapa := make(map[string]int)

	// Agregar elementos al mapa
	mapa["manzana"] = 5
	mapa["banana"] = 8
	mapa["cereza"] = 3

	fmt.Println("Mapa:", mapa)

	// Eliminar un elemento del mapa
	delete(mapa, "banana")
	fmt.Println("Mapa después de eliminar:", mapa)

	// Verificar si una clave existe en el mapa
	valor, ok := mapa["banana"]
	fmt.Println("Valor de banana:", valor, "¿Existe?", ok)

	// Recorrer el mapa
	for clave, valor := range mapa {
		fmt.Println("Clave:", clave, "Valor:", valor)
	}

	// Longitud del mapa
	fmt.Println("Longitud del mapa:", len(mapa))

	// Agregar un elemento al mapa si no existe desde una función
	agregarElemento(mapa, "banana", 10)
	fmt.Println("Mapa después de agregar:", mapa)

	// Función para agregar un elemento al mapa si no existe
	func agregarElemento(m map[string]int, clave string, valor int) {
		if _, ok := m[clave]; !ok { // Verificar si la clave no existe
			m[clave] = valor // Agregar elemento al mapa
		}
	}
}

/*________________________*/
/* Estructuras (Structs) */

type Persona struct {
	Nombre string
	Edad   int
	Ciudad string
}

func main() {
	// Inicialización de una estructura
	p := Persona{
		Nombre: "Juan",
		Edad:   30,
		Ciudad: "Madrid",
	}

	fmt.Println("Persona:", p)
}

/*________________________*/
/* Goroutines y Canales */

// Goroutine básica
func saludo(nombre string) {
	fmt.Println("Hola,", nombre)
}

func main() {
	// Ejecutar una goroutine
	go saludo("Go")

	// Esperar para que la goroutine se ejecute (simulación)
	var input string
	fmt.Scanln(&input)
}

/*________________________*/
/* Goroutines y Canales */

func sumar(a, b int, canal chan int) {
	canal <- a + b // Enviar el resultado al canal
}

func main() {
	canal := make(chan int) // Canal de enteros

	// Llamar a una goroutine
	go sumar(5, 3, canal)

	// Esperar el resultado desde el canal
	resultado := <-canal
	fmt.Println("Resultado de la suma:", resultado)
}

/*________________________*/
/* Interfaces */

type Animal interface {
	HacerSonido() string
}

type Perro struct {
	Nombre string
}

func (p Perro) HacerSonido() string {
	return "Guau!"
}

type Gato struct {
	Nombre string
}

func (g Gato) HacerSonido() string {
	return "Miau!"
}

func main() {
	var a Animal

	// Instanciación de objetos que implementan la interfaz Animal
	a = Perro{Nombre: "Rex"}
	fmt.Println(a.HacerSonido())

	a = Gato{Nombre: "Luna"}
	fmt.Println(a.HacerSonido())
}

/*________________________*/
/* Funciones como Primeros Ciudadanos */

func operar(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	suma := func(a, b int) int { return a + b }
	resta := func(a, b int) int { return a - b }

	fmt.Println("Suma:", operar(5, 3, suma))
	fmt.Println("Resta:", operar(5, 3, resta))
}

/*________________________*/
/* Función con Variadic Parameters */

func imprimirNumeros(numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(numero)
	}
}

func main() {
	imprimirNumeros(1, 2, 3, 4, 5)
}

/*________________________*/
/* Defer, Panic y Recover */

func dividir(a, b int) int {
	defer fmt.Println("Finalizando función dividir...") // Se ejecuta al final de la función

	if b == 0 {
		panic("Error: división por cero")
	}
	return a / b
}

func main() {
	// Usando recover para manejar panics
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de:", r)
		}
	}()

	fmt.Println("Resultado:", dividir(10, 2))
	fmt.Println("Resultado:", dividir(10, 0)) // Esto causa panic
}
