package main

import "fmt"

//Quiero hacer una condicional múltiple, haremos la de estatura y pesos para una montaña rusa

func main() {

	fmt.Println("Bienvenido a la montaña rusa Gaskull")

	//Variable de edad
	var edad int
	fmt.Println("¿Cuántos años tienes?: ")
	//Ingresa la respuesta y se corrobora error por tipo de ingreso
	_, err := fmt.Scan(&edad)
	if err != nil {
		fmt.Println("Chiko, debes ingresar un número")
		return
	}

	//Variable de pinchos requeridos
	var pinchos int
	fmt.Println("¿Cuántos pinchos tienes chiko?: ")
	_, err = fmt.Scan(&pinchos) //Se usa sólo el = para reasignar valor, el uso de := es para declarar
	if err != nil {
		fmt.Println("Chiko, debes ingresar un número")
		return
	}
	if edad >= 15 && pinchos >= 10 {
		fmt.Println("Puedes pasar chiko, disfruta el viaje.")
	} else if edad < 15 && pinchos >= 10 {
		fmt.Println("Aún no eres lo suficientemente verde.")
	} else if edad >= 15 && pinchos < 10 {
		fmt.Println("Ve a golpear a otro chiko y consigue más pinchos.")
	} else {
		fmt.Println("Vete de aquí humano, no eres verde ni tienes pinchos.")
	}
}
