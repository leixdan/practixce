package main

import "fmt"

func main() {

	fmt.Println("Bienvenidos")
	muestra := []int{100, 22, 31, 56, 90}
	fmt.Println(muestra)

	// Número mayor del slice
	mayor(muestra)

}

func mayor(slice []int) {
	max := 0
	for _, val := range slice {
		fmt.Println(val)
		if val > max {
			max = val
		}
	}
	fmt.Println(max)
}
