package main

import "fmt"

func main() {
	var x1 int
	fmt.Print("Qual o valor do seu número? ")
	fmt.Scan(&x1)
	dobro := x1 * 2
	triplo := x1 * 3
	quadruplo := x1 * 4
	fmt.Println("O dobro, triplo e quadruplo são respectivamente ", dobro, triplo, quadruplo)
}
