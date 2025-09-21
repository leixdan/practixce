package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Bienvenido al juego de adivinar")
	fmt.Println("y/n")
	var respuesta string
	fmt.Scanln(&respuesta)

	if respuesta == "y" {
		jugar()
	}

	fmt.Println("Adiós")

}

func jugar() {
	numRand := rand.Intn(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un número (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numRand {
			fmt.Println("Feliciddes, adivinaste el número.")
			return
		} else if numIngresado < numRand {
			fmt.Println("El número a adivinar es mayor")
		} else if numIngresado > numRand {
			fmt.Println("El número a adivinar es menor")
		}
	}
	fmt.Println("Se te acabaron los intentos, el número era: ", numRand)
}
