//Crea una función que reciba dos cadenas como parámetro (str1, str2)
//e imprima otras dos cadenas como salida (out1, out2).
//out1 contendrá todos los caracteres presentes en la str1 pero NO estén presentes en str2.
//out2 contendrá todos los caracteres presentes en la str2 pero NO estén presentes en str1.

package main

import (
	"fmt"
)

//var Sout1, Sout2 string

func convers(sOut1 []rune, sOut2 []rune, outArmado1 *string) {

	for _,r1 := range sOut1 {
		for _,r2 := range sOut2{

			if r1 == r2{
				*outArmado1 += string(r1)
				break
			}
	
		}
}
}

func main() {

	var str1, str2 string = "caramelo", "camaron"
	sOut1 := []rune(str1)
	sOut2 := []rune(str2)
	
	var outArmado1 string

	//fmt.Println("Ingresa las dos palabras")
	fmt.Println("Caramelo: ", sOut1)
	fmt.Println("Camaron: ",sOut2)

	convers(sOut1, sOut2, &outArmado1)
fmt.Println("_________________________")

fmt.Println(outArmado1)


}


/*

func convers(sOut1 []rune, sOut2 []rune) string {
	var resultado string
	for _, r1 := range sOut1 {
		for _, r2 := range sOut2 {
			if r1 == r2 {
				resultado += string(r1)
				break
			}
		}
	}
	return resultado
}

func main() {
	var str1, str2 string = "caramelo", "camaron"
	sOut1 := []rune(str1)
	sOut2 := []rune(str2)

	outArmado1 := convers(sOut1, sOut2)

	fmt.Println("_________________________")
	fmt.Println(outArmado1)
}
	*/

