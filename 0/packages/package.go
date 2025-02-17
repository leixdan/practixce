package ix

import "fmt"

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
