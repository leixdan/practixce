package main

import "fmt"

func main() {

	var gang [3]string // array of 3 strings

	gang[0] = "Nasus"
	gang[1] = "Queso"
	gang[2] = "Smol"

	fmt.Println(gang[2]) // print the third element

	// slice
	var ganga = []string{} // El mapa se declaran con llaves vacias

	ganga = append(ganga, "Tania") // element 0 is Piedrita
	ganga = append(ganga, "Ally") // element 1 is Monstera
	ganga = append(ganga, "Iran") // element 2 is Iran
	ganga = append(ganga, "Marciana") // element 3 is Marciana
	ganga = append(ganga, "Joana") // element 4 is Joana
	ganga = append(ganga, "Liz") // element 5 is Liz
	ganga = append(ganga, "Paula") // element  is Paula

	ganga = ganga[1:6] // remove the first element and the last element

	fmt.Println(ganga)

	
	// iterate over the slice
	for i,member := range ganga {
		fmt.Println("Index:", i+1, "Member:", member)		
	}
 
	
	// iterate over the slice without the index
	for _,member := range ganga { // _ is a blank identifier
		if member != "Marciana" { // skip the element
			continue // skip the element
		}	
		fmt.Println("Es un alien:", member)
	}
// convertir un string en slice y filtrar segun la ultima letra
	for _,member := range ganga {

		index :=  len(member)-1
		letter := member[index:]

		if letter == "a" { 
			continue
		}	
		fmt.Println("No termina en a:", member)
	}

	
}