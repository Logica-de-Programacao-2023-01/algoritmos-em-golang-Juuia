package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var i, j int

	fmt.Print("Digite o primeiro índice: ")
	fmt.Scan(&i)
	fmt.Print("Digite o segundo índice: ")
	fmt.Scan(&j)

	aux := slice[i]
	slice[i] = slice[j]
	slice[j] = aux

	fmt.Println("Slice resultante: ", slice)

}
