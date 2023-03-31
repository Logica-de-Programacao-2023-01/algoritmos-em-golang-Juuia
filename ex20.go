package main

import "fmt"

func main() {
	var tamanho, x int
	var slice []int

	fmt.Print("Qual o tamanho desse slice? ")
	fmt.Scan(&tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Print("Digite um número: ")
		fmt.Scan(&x)
		slice = append(slice, x)
	}

	crescente := true
	anterior := slice[0]
	for i := 1; i < len(slice); i++ {
		atual := slice[i]
		if atual < anterior {
			fmt.Println("Não está em ordem crescente")
			crescente = false
			break
		}
		anterior = atual
	}
	if crescente {
		fmt.Println("Está em ordem crescente")
	}
}
