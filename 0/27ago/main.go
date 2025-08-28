package main

import "fmt"

type Estudiante struct {
	Nombre         string
	Edad           int
	Calificaciones map[string]float32
}

func main() {

	fmt.Println("Bienvenido")

	e1 := Estudiante{
		Nombre:         "Rodrigo",
		Edad:           18,
		Calificaciones: make(map[string]float32),
	}
	e1.Calificaciones["Historia"] = 7.7

	fmt.Println(e1)

	e1.Calificaciones["Cálculo"] = 4.5
	e1.Calificaciones["Lógica"] = 10

	fmt.Println(e1)
}
