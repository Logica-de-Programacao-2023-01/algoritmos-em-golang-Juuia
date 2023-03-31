package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string

	fmt.Print("Digite sua palavra: ")
	fmt.Scan(&s)
	s = strings.ToUpper(s)
	fmt.Println("Sua palavra em maiúsculo é: ", s)
}
