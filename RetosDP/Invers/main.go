package main

import "fmt"

func main() {

	var solicitud string

	fmt.Println("Vamos a invertir el texto que me pases")
	fmt.Println("Ingresa la palabra que deseas invertir:")
	fmt.Scan(&solicitud)

	sSolicitud := []rune(solicitud)
	var invertido string

	for p := len(sSolicitud) - 1; p >= 0; p-- {
		fmt.Println(sSolicitud[p])
		invertido += string(sSolicitud[p])
	}

	fmt.Println(invertido)

}
