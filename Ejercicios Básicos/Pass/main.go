package main

import "fmt"

func main() {

	useradmin := "King"
	passadmin := "1234"

	var user string
	var pass string
	var log bool

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
			fmt.Println("Usuario incorrecto")
			break
		}
	}

	if log == true {
		fmt.Println("Bienvenido al sistema")
	} else {
		fmt.Println("Usuario o contraseña incorrectos")
	}
}
