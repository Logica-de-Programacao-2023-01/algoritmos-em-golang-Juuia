package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Gosto de comer bolo"
	if strings.Contains(s, "Banana") {

		fmt.Println("A string contém a substring 'Banana'")

	} else {

		fmt.Println("A string não contém a substring 'Banana'")

	}

}
