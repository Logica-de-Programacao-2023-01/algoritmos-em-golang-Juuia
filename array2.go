package main

import "fmt"

func main() {
	var nomes []string
	fmt.Println(nomes)

	nomes = append(nomes, "João")
	fmt.Println(nomes)

	nomes = append(nomes, "Maria")
	fmt.Println(nomes)

	nomes = append(nomes, "José")
	fmt.Println(nomes)
}
