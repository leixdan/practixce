package main

import "fmt"

func main() {
//Usuario y contraseña a comparar
	useradmin := "King"
	passadmin := "1234"

//Variables a usar y leer para el inicio de sesión
	var user string
	var pass string
	var log bool
	//var try int

//Ciclo del login
	for {

		fmt.Println("¿Cuál es tu usuario?: ")
		fmt.Scanln(&user)
		fmt.Println("¿Cuál es tu contraseña?: ")
		fmt.Scanln(&pass)

		if user == useradmin && pass == passadmin {
			log = true
			break
		} else if user == useradmin && pass != passadmin {
			fmt.Println("Contraseña incorrecta")
			log = false
			break
		} else {
			break
		}
	}

	if log == true {
		fmt.Println("Bienvenido al sistema")
	} else {
		fmt.Println("Usuario o contraseña incorrectos")
	}
}
