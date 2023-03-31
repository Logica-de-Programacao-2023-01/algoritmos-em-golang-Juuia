package main

import "fmt"

func main() {
	var primeira, segunda string

	fmt.Print("Qual a sua primeira frase? ")
	fmt.Scan(&primeira)
	fmt.Print("Qual a sua segunda frase? ")
	fmt.Scan(&segunda)

	s3 := primeira + ", " + segunda
	fmt.Println(s3)
}
