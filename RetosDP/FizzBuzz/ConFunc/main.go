package main

import "fmt"

/*
for i := 1; i <= 10; i++ {
		if esMultiploDeCinco(i) {
			fmt.Printf("%d es Múltiplo de 5\n", i)
			continue
		} else {
			fmt.Println(i)
		}
	}
*/

// Funciónes

// esMultiploDeCinco determina si un número es múltiplo de 5 y devuelve true o false
func esMultiploDeTres(n int) bool {
	return n%3 == 0
}

// esMultiploDeTres determina si un número es múltiplo de 3 y devuelve true o false
func esMultiploDeCinco(n int) bool {
	return n%5 == 0
}

func main() {
	for i := 1; i <= 100; i++ {
		if esMultiploDeTres(i) {
			fmt.Printf("%d Fizz\n", i)
			continue
		} else if esMultiploDeCinco(i) {
			fmt.Printf("%d Buzz\n", i)
			continue
		} else if esMultiploDeTres(i) && esMultiploDeCinco(i) {
			fmt.Printf("%d FizzBuzz\n", i)
			continue
		} else {
			fmt.Println(i)

		}
	}
}
