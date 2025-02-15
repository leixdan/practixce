//Gorro seleccionador gaskull 2.0

package main

import "fmt"

func lista() {

}

//Voy a necesitar las variables: "nombre", "edad", "r1", r2", "r3", "r4" y "resultado"

func main() {

	fmt.Println("Bienvenido al sombrero seleccionador")
	var nombre string
	fmt.Println("¿Cuál es tu nombre?: ")
	fmt.Scan(&nombre)
	fmt.Printf("Un gusto %s, ¿Cuántos años tienes?: ", nombre)
	var edad int
	fmt.Scan(&edad)
	fmt.Printf("Ya veo, así que tienes %d años, bien bien.\n", edad)
	// Hasta aquí se hacen las primeras preguntas
	// Inicialización del ciclo de respuestas
	for {

		//Declaración del mapa donde se guarden las respuestas

		respuestas := make(map[string]int)

		fmt.Println("Esta es la pregunta 1: ")
		var r1 int
		fmt.Scan(&r1)
		respuestas["r1"] = r1

		fmt.Println("Esta es la pregunta 2: ")
		var r2 int
		fmt.Scan(&r2)
		respuestas["r2"] = r2

		fmt.Println("Esta es la pregunta 3: ")
		var r3 int
		fmt.Scan(&r3)
		respuestas["r3"] = r3

		fmt.Println("Esta es la pregunta 4: ")
		var r4 int
		fmt.Scan(&r4)
		respuestas["r4"] = r4
		/*
		   fmt.Println("Tus respuestas fueron:")

		   	for clave,valor := range respuestas{
		   		fmt.Println(clave ,valor)
		   	}
		*/
		valor := respuestas["r1"]
		fmt.Printf("La respuesta a la pregunta 1 fue %d\n", valor)

		valor = respuestas["r2"]
		fmt.Printf("La respuesta a la pregunta 2 fue %d\n", valor)

		valor = respuestas["r3"]
		fmt.Printf("La respuesta a la pregunta 3 fue %d\n", valor)

		valor = respuestas["r4"]
		fmt.Printf("La respuesta a la pregunta 4 fue %d\n", valor)
		// Pregunta de confirmación
		fmt.Println("¿Estás de acuerdo con tus respuestas?: y/n")
		var conforme string
		fmt.Scan(&conforme)

		if conforme == "n" {
			fmt.Println("Ya veo, comencemos de nuevo...")
			continue
		}else if conforme == "y"{
			break
		}else{
			fmt.Println("Esa es una respuesta incorrecta, por favor ingresa (y) o (n)")
		}

	}
	fmt.Printf("Bien %s, podemos continuar\n", nombre)
}
