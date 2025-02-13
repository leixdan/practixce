//Programa en el que se guarden 3 ID con sus respectivos datos de los usuarios

package main

import "fmt"

//Primero se genera la struct

type solicitante struct {
	Nombre string
	Edad   int
}

func main() {
	// Se genra el mapa
	solicitantes := make(map[int]solicitante)
	// Datos del mapa con ID como clave y como dato los struct de amigos
	solicitantes[1] = solicitante{Nombre: "Alondra", Edad: 24}
	solicitantes[2] = solicitante{Nombre: "Iran", Edad: 26}
	solicitantes[3] = solicitante{Nombre: "Lizbeth", Edad: 27}

	
	fmt.Println("Bienvenido a la lista de amigos chiclebomba")
	var id int
	fmt.Println("Ingresa el ID del amigo a consultar: ")
	fmt.Scan(&id)

	switch id {
	case 1:
		amigo := solicitantes[1]
		fmt.Println("El amigo es:", amigo)
	case 2:
		amigo := solicitantes[2]
		fmt.Println("El amigo es:", amigo)
	case 3:
		amigo := solicitantes[2]
		fmt.Println("El amigo es:", amigo)
	
	}

}
