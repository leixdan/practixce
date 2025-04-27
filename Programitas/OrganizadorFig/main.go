package main

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

//Función que recibe una figura y devuelve su área
func ImprimirArea (f Figura) float64 { //Se usa la interfaz Figura como tipo de parámetro
	// Se puede pasar cualquier figura que implemente el método Area
	return f.Area()
}

func main() {





}
