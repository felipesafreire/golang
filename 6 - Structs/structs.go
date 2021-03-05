package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {

	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Felipe Sá Freire"
	u.idade = 12
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua A", 21}

	u2 := usuario{"Felipe", 15, enderecoExemplo}
	fmt.Println(u2)

	usuario3 := usuario{idade: 25}
	fmt.Println(usuario3)

}
