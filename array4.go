package main

import "fmt"

func main() {
	arr := [6]float64{5.2, 7.8, 2.6, 5.3, 4.7, 8.2}
	var soma float64

	for _, x := range arr {
		soma += x
	}

	media := soma / float64(len(arr))

	fmt.Println("A média é de", media)
}
