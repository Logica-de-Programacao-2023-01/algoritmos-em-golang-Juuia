package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var n int

	fmt.Print("Digite sua palavra: ")
	fmt.Scan(&s)
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	if n > len(s) {
		n = len(s)
	}
	s = strings.ToUpper(s[:n]) + s[n:]
	fmt.Println("A palavra com a(s) n primeiras letras em maiúsculo: ", s)
}
