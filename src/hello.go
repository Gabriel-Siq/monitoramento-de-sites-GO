package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

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

func iniciaMonitoramento() {

	tempo := 5

	fmt.Print("\n")
	fmt.Println("Monitorando...")

	sites := []string{"https://www.google.com.br", "https://www.youtube.com.br", "https://www.instagram.com/"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site", i, ":", site)
			time.Sleep(1 * time.Second)
			testaSite(site)
		}
		for i := 0; i < delay; i++ {
			fmt.Println("Um novo monitoramento comecará em ", tempo, "segundos")
			time.Sleep(1 * time.Second)
			tempo--
		}
		fmt.Println(("---------------------------"))
		tempo = delay
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro ao testar o site", site, ":", err)
		return
	}

	if resp != nil && resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado sem erros!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

	fmt.Println(("---------------------------"))
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, _ := os.Open("sites.txt")
	fmt.Println(arquivo)

	return sites
}

func main() {

	Introduz()
	leSitesDoArquivo()

	for {
		escolha := Escolhe()

		switch escolha {
		case 1:
			iniciaMonitoramento()
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
}
