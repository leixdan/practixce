//Crea una función que reciba dos cadenas como parámetro (str1, str2)
//e imprima otras dos cadenas como salida (out1, out2).
//out1 contendrá todos los caracteres presentes en la str1 pero NO estén presentes en str2.
//out2 contendrá todos los caracteres presentes en la str2 pero NO estén presentes en str1.

package main

import "fmt"

func convers(sOut1 []rune, sOut2 []rune) string {	//La función recibe dos runes y devuelve un string
	var resultado string							//Se declara una variable local de la función
	for _, r1 := range sOut1 {						//Se recorre r1(rune1)
		for _, r2 := range sOut2 {					//Ciclo anidado de r2(rune2)
			if r1 == r2 {							//Se compara el recorrido de r1 y r2
				resultado += string(r1)				//Se guarda en la variable temporal resultado si son iguales
				fmt.Println(resultado)
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


