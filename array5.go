package main

import "fmt"

func main() {
	var matrix [3][4]int

	for linha := range matrix {
		for coluna := range matrix[linha] {
			matrix[linha][coluna] = linha + coluna
		}
	}

	fmt.Println(matrix)
}
