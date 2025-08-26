package main

import "fmt"

type Mascota struct {
	Nombre  string
	Edad    int
	Especie string
}

func cumpleaños(m *Mascota) {
	m.Edad++
}

func main() {

	nasus := Mascota{
		Nombre:  "Nasus",
		Edad:    5,
		Especie: "Perro",
	}

	fmt.Printf("El nombre de la mascota es: %s, tiene %d años y es un %s.\n", nasus.Nombre, nasus.Edad, nasus.Especie)

	cumpleaños(&nasus)

	fmt.Printf("Hoy cumplió años y ahora tiene: %d años\n", nasus.Edad)
}
