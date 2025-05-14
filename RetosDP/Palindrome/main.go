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

	for i, l := 0, len(slice_palabra)-1; i < l; i, l = i+1, l-1 {
		fmt.Printf("i: %d, l:%d\n", slice_palabra[i], slice_palabra[l])
	}

}
