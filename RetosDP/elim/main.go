//Crea una función que reciba dos cadenas como parámetro (str1, str2)
//e imprima otras dos cadenas como salida (out1, out2).
//out1 contendrá todos los caracteres presentes en la str1 pero NO estén presentes en str2.
//out2 contendrá todos los caracteres presentes en la str2 pero NO estén presentes en str1.

/*
   for i := range a {
       if a[i] != b[i] {
           return false
       }
*/

package main

import (
	"fmt"
)

var Out1, Out2 string

func convers(sOut1 []rune, sOut2 []rune) { //(string, string) {

	for i := range sOut1 {
		fmt.Println(sOut1[i])

		if sOut1[i] == sOut2[i] {
			Out1 += string(sOut1[i])
		}
	}
	fmt.Println("Fin de ciclo 1")
	for i := range sOut2 {
		fmt.Println(sOut2[i])
	}

	//out1 := string(r1)
	//out2 := string(r2)
}

func main() {

	var str1, str2 string = "hada", "miada"
	sOut1 := []rune(str1)
	sOut2 := []rune(str2)

	fmt.Println("Ingresa las dos palabras")

	convers(sOut1, sOut2)

	fmt.Println(sOut1)
	fmt.Println(sOut2)
	fmt.Println("_________________________________")
	fmt.Println(Out1)

}
