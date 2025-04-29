package main

import (
	"fmt"
	preguntas "sombrero_seleccionador/Preguntas"
)

func main() {

	//Variables
	var ready, name string
	var respuesta string
	var puntaje preguntas.Puntaje

	fmt.Println("Hola, bienvenida o bienvenido al sombrero seleccionador")
	fmt.Println("A través de un breve quiz lograremos saber a que casa perteneces")
	fmt.Println("¿Estás lista o listo? (y/n)")
	fmt.Scan(&ready)

	if ready == "n" || ready == "N" {
		fmt.Println("Decidieste no continuar, bye")
	} else {
		println("Bien, ¿cuál es tu nombre? ")
		fmt.Scan(&name)
		fmt.Printf("Entonces te llamas %s, bien, ahora las preguntas, selecciona sólo la letra que elijas\n", name)

		//Preguntas

		fmt.Println("¿Cuál de estas cualidades valoras más en ti?: ")
		fmt.Println("a) Valentía")
		fmt.Println("b) Inteligencia")
		fmt.Println("c) Lealtad")
		fmt.Println("d) Ambición")
		fmt.Scan(&respuesta)
		preguntas.SumarRespuesta(&puntaje, &respuesta)

		fmt.Println("En un conflicto, tu primera reacción suele ser: ")
		fmt.Println("a) Enfrentarlo directamente")
		fmt.Println("b) Analizar la situación antes de actuar")
		fmt.Println("c) Buscar una solución pacífica")
		fmt.Println("d) Usarlo a tu favor")
		fmt.Scan(&respuesta)
		preguntas.SumarRespuesta(&puntaje, &respuesta)

		fmt.Println("¿Qué tipo de personas prefieres tener cerca?: ")
		fmt.Println("a) Valientes y decididas")
		fmt.Println("b) Curiosas y creativas")
		fmt.Println("c) Amables y confiables")
		fmt.Println("d) Motivadas y ambiciosas")
		fmt.Scan(&respuesta)
		preguntas.SumarRespuesta(&puntaje, &respuesta)

		fmt.Println("Tu lugar ideal para estudiar o trabajar sería: ")
		fmt.Println("a) Un lugar lleno de acción, aunque ruidoso")
		fmt.Println("b) Una biblioteca silenciosa")
		fmt.Println("c) Un rincón acogedor con amigos")
		fmt.Println("d) Una oficina moderna con vista al éxito")
		fmt.Scan(&respuesta)
		preguntas.SumarRespuesta(&puntaje, &respuesta)

		fmt.Println("Si ganaras un premio importante, ¿cómo reaccionarías?: ")
		fmt.Println("a) Lo celebrarías con entusiasmo y orgullo")
		fmt.Println("b) Te sentirías honrado, pero seguirías con tus estudios")
		fmt.Println("c) Lo compartirías con quienes te ayudaron")
		fmt.Println("d) Lo usarías como impulso para llegar aún más lejos")
		fmt.Scan(&respuesta)
		preguntas.SumarRespuesta(&puntaje, &respuesta)
	}
	fmt.Println(puntaje.A, puntaje.B, puntaje.C, puntaje.D)
}
