package main

import "fmt"

//Primero definimos la estructura que será base para la interfáz y los métodos

type pokemon struct {
	Nombre string
	Tipo   string
	Altura float32
}

type digimon struct {
	Nombre string
	Tipo   string
	Altura float32
}

// Método de Pokemon
func (p pokemon) describir() {
	fmt.Printf("Veamos los datos del pokemon elegido:\n")
	fmt.Printf("Nombre: %s\nTipo: %s\nAltura: %f\n", p.Nombre, p.Tipo, p.Altura)
}

func main() {

	var choose int

	garchomp := pokemon{
		Nombre: "Garchomp",
		Tipo:   "Dragón / Tierra",
		Altura: 1.90,
	}

	togekiss := pokemon{
		Nombre: "Togekiss",
		Tipo:   "Hada / Volador",
		Altura: 1.50,
	}

	fmt.Println("Ingresa el número del poke a revisar:")
	fmt.Println("1. Garchomp")
	fmt.Println("2. Togekiss")
	fmt.Scan(&choose)

	switch choose {
	case 1:
		garchomp.describir()
	case 2:
		togekiss.describir()
	}

}
