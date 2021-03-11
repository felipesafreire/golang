package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {

	totalDaSoma := soma(1, 2, 41, 15, 5, 8)
	fmt.Println(totalDaSoma)

	escrever("Olá Mundo", 2, 5, 7, 9, 4)

}
