package bendis

import "fmt"

var collection []string //Slice de los nombres

func Add(name string) {
	collection = append(collection, name) //Se añade el valor en "name"
}

func list() {
	for i, n := range collection {
		fmt.Printf("%d, %s\n", i+1, n)
	}
}

/*
//Las variables que inicien con minuscula son privadas
var lottus = "Education"
//Si está en el mismo paquete se puede acceder a la variable privada


//Las variables que inicien con mayuscula son publicas"
var Lottus = "Quality"

//Las funciones que inicien con minuscula son privadas
func lottusFunc() {
	fmt.Println("Education")
}

//Las funciones que inicien con mayuscula son publicas
func LottusFunc() {
	fmt.Println("Quality")
}
*/
