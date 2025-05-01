//Funciones para el programa

package funciones

import (
	"citasmed/agenda"
)

// Init inicializa el mapa en agenda.go
func Init() {
	agenda.Agenda = make(map[string]agenda.Beneficiario)
}
