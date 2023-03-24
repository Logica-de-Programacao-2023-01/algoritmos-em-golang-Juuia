package main

import (
	"fmt"
)

func main() {
	var x, y, z int
	fmt.Print("Digite o primeiro numero: ")
	fmt.Scan(&x)
	fmt.Print("Digite o segundo numero: ")
	fmt.Scan(&y)
	fmt.Print("Digite o terceiro numero: ")
	fmt.Scan(&z)

	if x > y && x > z {
		if y > z {
			fmt.Printf("%d > %d > %d", x, y, z)
		} else {
			fmt.Printf("%d > %d > %d", x, y, z)
		}
	} else if y > x && y > z {
		if x > z {
			fmt.Printf("%d > %d > %d", x, y, z)
		} else {
			fmt.Printf("%d > %d > %d", x, y, z)
		}
	} else {
		if x > y {
			fmt.Printf("%d > %d > %d", z, x, z)
		} else {
			fmt.Printf("%d > %d > %d", z, y, x)
		}

	}

}
