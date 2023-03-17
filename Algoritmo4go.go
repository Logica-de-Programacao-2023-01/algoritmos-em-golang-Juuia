package main

import "fmt"

func main() {
	var x1, x2, x3 float64
	fmt.Print("Qual o valor do seu primeiro número? ")
	fmt.Scan(&x1)
	fmt.Print("Qual o valor do seu segundo número? ")
	fmt.Scan(&x2)
	fmt.Print("Qual o valor do seu terceiro número? ")
	fmt.Scan(&x3)

	resultado := ((x1 * 2) + (x2 * 3) + (x3 * 5)) / 10
	fmt.Println("Sua média ponderada é ", resultado)

}
