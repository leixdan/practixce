package preguntas

//Es mejor hacerlo en un struct perrito

type Puntaje struct {
	A int
	B int
	C int
	D int
}

// SumarRespuesta sumará en la estructura Puntaje
func SumarRespuesta(p *Puntaje, respuesta *string) {
	switch *respuesta {
	case "a", "A":
		p.A++
	case "b", "B":
		p.B++
	case "c", "C":
		p.C++
	case "d", "D":
		p.D++
	}
}
