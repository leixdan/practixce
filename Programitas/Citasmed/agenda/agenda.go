//Aquí vamos a guardar los types de las citas para el día

package agenda

//Consultorios usa arreglos ya que no cambiarán

var Consultorios = [11]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

//Datos de las personas con cita

type Adress struct {
	Calle      string
	Colonia    string
	Delegación string
}

type Beneficiario struct {
	Nombre      string
	Edad        int
	Consultorio int
	Horario     string
	Domicilio   Adress
}

var Agenda map[string]Beneficiario
