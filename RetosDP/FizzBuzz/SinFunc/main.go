//Escribir un programa que en consola muestre números del 1 al 100, sustituyendo:
// los múltiplos de 3 por la palabra “Fizz”
// los múltiplos de 5 por la palabra “Buzz”
// los múltiplos de 3 y 5 por la palabra “FizzBuzz”

/*
for i := 1; i <= 10; i++ {
		if i%5 == 0 {
			fmt.Printf("%d Es múiltiplo de 5\n", i)
			continue
		}
		fmt.Println(i)
	}
}

*/

package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Printf("%d Fizz\n", i)
			continue
		}
		if i%5 == 0 {
			fmt.Printf("%d Buzz\n", i)
			continue
		}
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d FizzBuzz\n", i)
			continue
		}
		fmt.Println(i)
	}

}
