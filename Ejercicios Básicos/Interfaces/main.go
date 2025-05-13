//Catálogo multimedia
//Mezcla de libros, películas y podcasts.
//Cada uno tiene distintos metadatos (autor, duración, género).
//Necesitas crear una vista ordenada por título o tipo de contenido.

package main

import "fmt"

// Primero definimos la estructura que será base para la interfáz y los métodos
type books struct {
	Nombre    string
	Editorial string
	Autor     string
	Año       int
}

type movies struct {
	Nombre   string
	Director string
	Género   string
	Año      int
}

type photos struct {
	Nombre string
	Estilo string
	Autor  string
	Año    int
}

// Métodos asociados
func (m movies) Sound() {
	fmt.Println("Este archivo multimedia tiene sonido")
}

func (b books) Pages() {
	fmt.Println("Este archivo multimedia tiene páginas")
}

// Interfaces
type Audio interface {
	Sound() string
}

type Page interface {
	Pages() string
}

func main() {

	cl := books{
		Nombre:    "Cuervo y Luna",
		Editorial: "Atipica Espectral",
		Autor:     "Katherine",
		Año:       2025,
	}

	mm := movies{
		Nombre:   "Mad max",
		Director: "Nasus",
		Género:   "Acción",
		Año:      2021,
	}

	cl.Pages()
	mm.Sound()

}
