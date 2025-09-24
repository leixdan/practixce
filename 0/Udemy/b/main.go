package main

import "fmt"

func main() {

	fmt.Println("Bienvenidos")
	muestra := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(muestra)

	// Número mayor del slice
	mayor(muestra)

}

func mayor(slice []int) {
	max := 0
	fmt.Println(max)
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
}
