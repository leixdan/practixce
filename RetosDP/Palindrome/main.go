package main

import "fmt"

/*func revisar(palabra *string) {

	for i, l := range *palabra {
		fmt.Printf("Índice: %d Letra: %c\n", i, l)
	}
}*/

func main() {

	var palabra string

	fmt.Println("Vamos a ver si tu palabra es un palíndromo")
	fmt.Println("Ingresa tu palabra")
	fmt.Scan(&palabra)
	//revisar(&palabra)

	slice_palabra := []rune(palabra)
	esPalindromo := true

	for i, l := 0, len(slice_palabra)-1; i < l; i, l = i+1, l-1 {
		if slice_palabra[i] != slice_palabra[l] {
			esPalindromo = false
			break
		}
	}

	if esPalindromo {
		fmt.Println("¡Sí es un palíndromo!")
	} else {
		fmt.Println("No es un palíndromo.")
	}

}
