package main

import "fmt"

func main() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var buscado int

	fmt.Print("Digite um número: ")
	fmt.Scan(&buscado)

	var encontrado bool

	for _, x := range arr {
		if x == buscado {
			fmt.Println("O elemento está dentro do array")
			encontrado = true
			break
		}
	}
	if !encontrado {
		fmt.Println("Não encontrei")
	}
}
