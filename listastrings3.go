package main

import (
	"fmt"
)

func main() {
	var x, c string
	var y int

	fmt.Print("Digite uma palavra ")
	fmt.Scan(&x)
	fmt.Print("Digite seu caractere: ")
	fmt.Scan(&c)

	for _, i := range x {
		if string(i) == c {
			y++
		}
	}
	fmt.Printf("O caractere %s ocorre %d vezes na string %s.\n", c, y, x)
}
