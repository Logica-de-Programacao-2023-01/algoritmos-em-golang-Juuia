package main

import "fmt"

func main() {
	var palavra string

	fmt.Print("Digite uma palavra ")
	fmt.Scan(&palavra)

	s := palavra
	fmt.Println("len(s) =", len(s))

}
