//Funciones para el programa

package funciones

import (
	"citasmed/agenda"
	"fmt"
)

// Init inicializa el mapa en agenda.go
func Init() {
	agenda.Agenda = make(map[string]agenda.Beneficiario)
}

func RevisarConsultorios (){
	for _, c := range agenda.Agenda{
		fmt.Printf("El consultorio de: %s es el número: %d\n", c.Nombre ,c.Consultorio)
	}
}