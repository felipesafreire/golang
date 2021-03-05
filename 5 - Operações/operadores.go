package main

import "fmt"

func main() {

	// ARITMETICOS

	soma := 1 + 1
	subtracao := 1 - 1
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 5

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25

	var soma2 int16 = numero1 + numero2
	fmt.Println(soma2)

	// FIM ARITMETICOS

	// ATRIBUIÇÃO

	var variavel1 string = "String 1"
	variavel2 := "String 2"

	fmt.Println(variavel1, variavel2)

	// FIM ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// FIM OPERADORES RELACIONAIS

	// OPERADORES LÓGICOS
	fmt.Println("---------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// FIM OPERADORES LÓGICOS

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15 //numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 //numero = numero - 20
	fmt.Println(numero)

	numero *= 3 //numero = numero * 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	// FIM OPERADORES UNÁRIOS

	var text string
	if numero > 5 {
		text = "Maior que 5"
	} else {
		text = "Menor que 5"
	}

	fmt.Println(text)

}
