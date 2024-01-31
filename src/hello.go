package main

import (
	"fmt"
	"os"
)

func Introduz() {
	var nome string = "Gabriel"

	fmt.Print("Seja bem-vindo Sr.", nome)
	fmt.Print("\n")
}

func Escolhe() int {
	var comando int
	fmt.Println("Escolha sua opção:")
	fmt.Print("1 - Monitoramento dos Sites")
	fmt.Print("\n")
	fmt.Print("2 - Análise dos Logs")
	fmt.Print("\n")
	fmt.Print("3 - Sair")
	fmt.Print("\n")
	fmt.Print("Escolha: ")
	fmt.Scanf("%d", &comando)

	return comando
}

func main() {

	Introduz()
	escolha := Escolhe()

	switch escolha {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo os Logs...")
	case 3:
		fmt.Println("Saindo do Programa...")
		os.Exit(0)
	default:
		fmt.Println("Sua escolha não é válida")
		os.Exit(-1)
	}
}
