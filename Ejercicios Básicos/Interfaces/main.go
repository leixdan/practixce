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
func (m movies) Sound() string {
	return "Este archivo multimedia tiene sonido"
}
func (b books) Pages() string {
	return "Este archivo multimedia tiene páginas"
}
func (p photos) Images() string {
	return "Este archivo multimedia tiene imágenes"
}

// Interfaces
type Sounds interface {
	Sound() string
}

func itSounds(s Sounds) {
	fmt.Println(s.Sound())
}

type Page interface {
	Pages() string
}

func itPages(p Page) {
	fmt.Println(p.Pages())
}

type Image interface {
	Images() string
}

func itImages(i Image) {
	fmt.Println(i.Images())
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

	mc := photos{
		Nombre: "Mi cielo",
		Estilo: "B/N",
		Autor:  "Jalapeño",
		Año:    2023,
	}

	fmt.Println("Datos por interfaz")
	itSounds(mm)
	itPages(cl)
	itImages(mc)
}
