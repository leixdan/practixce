//Este programa usará los otros archivos a modo de módulos para poder correr
//Se deberá tener el menú de acceso a los módulos: 1. Consulta de citas, 2. Lista de consultorios, 3. Personas en espera y 4. Salir
//En 1 deberá dar las opciones de eliminar una cita programada y colocar a una persona en espera, deberá descontar la persona en espera

package main

import (
	"citasmed/agenda"
	"citasmed/funciones"
	"fmt"
)

func main() {

	fmt.Println("Bienvenido al consultorio médico")
	funciones.Init()

	MX6 := agenda.Beneficiario{
		Nombre:      "Queso",
		Edad:        3,
		Consultorio: 5,
		Horario:     "M",
		Domicilio: agenda.Adress{
			Calle:      "Emiliano Zapata 67",
			Colonia:    "Azteca",
			Delegación: "Venustiano Carranza",
		},
	}
	//Si no declaro "MX6" como clave para el mapa, no furula
	agenda.Agenda["MX6"] = MX6

	for clave, beneficiario := range agenda.Agenda {
		fmt.Printf("Clave: %s\nDatos:%+v\n", clave, beneficiario)
	}

	funciones.RevisarConsultorios()

}
