package main

import "fmt"

func main() {
	var soma, tamanho, x int
	var slice []int

	fmt.Print("Qual o tamanho desse slice? ")
	fmt.Scan(&tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)
		slice = append(slice, x)
		soma += x
	}

	fmt.Println("O slice: ", slice)
	fmt.Println("A soma é de: ", soma)
}
