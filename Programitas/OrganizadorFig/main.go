package main

import "fmt"

type Figura interface { // Toda figura tiene un area
	Area() float64
}

// Rectángulo
type Rectangulo struct {
	Ancho, Alto float64
}

func (r Rectangulo) Area() float64 { // Implementación del método Area para Rectángulo
	return r.Ancho * r.Alto
}

// Círculo
type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 { // Implementación del método Area para Círculo
	return 3.14 * c.Radio * c.Radio
}

// Triángulo
type Triangulo struct{
	Base, Alto float64
}
func (t Triangulo) Area() float64{
	return t.Base * t.Alto / 2
}

// Función que recibe una figura y devuelve su área
func ImprimirArea(f Figura) { //Se usa la interfaz Figura como tipo de parámetro
	// Se puede pasar cualquier figura que implemente el método Area
	fmt.Printf("El área de la figura es: %.2f\n", f.Area())
}

func main() {

	r := Rectangulo{Ancho: 10, Alto: 5}
	c := Circulo{Radio: 3}
	t := Triangulo{Base: 9, Alto: 10}

	ImprimirArea(r)
	ImprimirArea(c)
	ImprimirArea(t)

}
