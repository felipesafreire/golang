package main

import (
	"errors"
	"fmt"
)

func main() {

	// NUMEROS INTEIROS

	var numero int64 = 100000000
	fmt.Println(numero)

	var numero2 uint32 = 100
	fmt.Println(numero2)

	//alias
	// INT32 = Rune
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = uint8
	var numero4 byte = 132
	fmt.Println(numero4)

	// FIM NUMEROS INTEIROS

	//NUMEROS REAIS

	var numeroReal1 float32 = 1200003.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1200000000000003.45
	fmt.Println(numeroReal2)

	numeroReal3 := 123.45
	fmt.Println(numeroReal3)

	// FIM NUMEROS REAIS

	// STRINGS

	var str1 string = "Texto"
	fmt.Println(str1)

	str2 := "Texto2"
	fmt.Println(str2)

	char := '3'
	fmt.Println(char)

	// FIM STRINGS

	// VARIAVEL SEM DADOS

	var texto string
	fmt.Println(texto)

	// FIM VARIAVEL SEM DADOS

	// BOOL

	var booleano1 bool = true
	fmt.Println(booleano1)

	var booleano2 bool = false
	fmt.Println(booleano2)

	var booleano3 bool
	fmt.Println(booleano3)

	// FIM BOOL

	// TIPO ERROR

	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)

	// FIM TIPO ERROR

}
