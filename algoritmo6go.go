package main

import "fmt"

func main() {

	var salario float64
	fmt.Print("Qual o valor do seu salário? ")
	fmt.Scan(&salario)

	resultado := ((salario / 100) * 15) + salario
	fmt.Println("seu novo salário com aumento de 15% é ", resultado)

}
