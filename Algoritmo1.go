package main

import "fmt"

func main() {
	var x1, x2, x3 int
	fmt.Print("Qual o valor do seu primeiro número? ")
	fmt.Scan(&x1)
	fmt.Print("Qual o valor do seu segundo número? ")
	fmt.Scan(&x2)
	fmt.Print("Qual o valor do seu terceiro número? ")
	fmt.Scan(&x3)
	soma := x1 + x2 + x3
	fmt.Println("a soma é ", soma)
}
