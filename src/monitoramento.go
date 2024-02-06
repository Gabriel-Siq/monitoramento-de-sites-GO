package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando o site", (i + 1), ":", site)
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
		registraLogs(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLogs(site, false)
	}

	fmt.Println(("---------------------------"))
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)
		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registraLogs(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}

func main() {

	Introduz()

	for {
		escolha := Escolhe()

		switch escolha {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo os Logs...")
			imprimeLogs()
		case 3:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Sua escolha não é válida")
			os.Exit(-1)
		}
	}
}
