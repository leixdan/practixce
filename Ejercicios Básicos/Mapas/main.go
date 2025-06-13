package main

import "fmt"

// //////////////////////////////////////////////////////////////////////////////
func elim(str string) string {
	m := make(map[rune]bool)
	var res []rune

	for _, r := range str {
		m[r] = true
		fmt.Println(m)
		res = append(res, r)
	}

	return string(res)

}

func main() {

	s1 := "Mondongo"

	r1 := elim(s1)
	fmt.Println(r1)

}
