package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	aula5()
}

func aula1() {
	fmt.Println("aula 1")
	fmt.Println("------")
	fmt.Println("Hello World")
}
func aula2() {
	fmt.Println("aula 2")
	fmt.Println("------")
	var nome string = "Douglas"
	var versao float32 = 1.1
	var idade int = 20
	fmt.Println("Olá Sr.", nome, "e sua idade é ", idade)
	fmt.Println("Este programa esta na versão", versao)
}
func aula3() {
	fmt.Println("aula 3")
	fmt.Println("------")
	nome := "Douglas"
	versao := 1.1
	idade := 24
	fmt.Println("Olá Sr.", nome, "e sua idade é ", idade)
	fmt.Println("Este programa esta na versão", versao)
	fmt.Println("O tipo da variavel nome é ", reflect.TypeOf(nome))
}

func aula4() {
	fmt.Println("aula 4")
	fmt.Println("------")
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O endereco da minha variavel comando é ", &comando)
	fmt.Println("O comando escolhido foi ", comando)
	var comando2 int
	//ele infere o tipo da variavel
	fmt.Scan(&comando2)
	fmt.Println("O endereco da minha variavel comando é ", &comando2)
	fmt.Println("O comando escolhido foi ", comando2)
}
func aula5() {
	fmt.Println("aula 5")
	fmt.Println("------")
	var nome string
	versao := 1.1
	fmt.Println("Digite seu nome")
	fmt.Scan(&nome)
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	arquivo, err := os.Open("file.txt")
	sites := []string{}

	if err != nil {
		fmt.Println(err)
	} else {
		leitor := bufio.NewReader(arquivo)
		for {
			linha, err2 := leitor.ReadString('\n')
			if err2 == nil {
				sites = append(sites, linha)
			} else {
				break
			}
		}
	}

	for {
		comando := mostrarMenu()

		switch comando {
		case 0:
			iniciarMonitoramento(sites)
		case 1:
			exibirLogs()
		case 2:
			saindoDoPrograma()
		default:
			comandoNaoReconhecido()
		}
	}
}
func mostrarMenu() int {
	fmt.Println("0- Iniciar Monitoramento")
	fmt.Println("1- Exibir Logs")
	fmt.Println("2- Sair do Programa")
	var comando int
	fmt.Scanf("%d", &comando)
	return comando
}

func iniciarMonitoramento(sites []string) {
	fmt.Println("Iniciando Monitoramento...")

	for {
		fmt.Println("Testes realizados em ", time.Now().Format("01-02-2006 15:04:05"))
		for _, site := range sites {
			resp, err := http.Get(strings.TrimSpace(site))
			fmt.Println(verificaStatus(resp, err, site))
		}

		time.Sleep(1 * time.Minute)
	}

}
func exibirLogs() {
	fmt.Println("Exibindo logs...")
	mostrarLogs()
}
func saindoDoPrograma() {
	fmt.Println("Saindo do programa...")
	os.Exit(0)
}
func comandoNaoReconhecido() {
	fmt.Println("Comando não reconhecido")
	os.Exit(-1)
}

func verificaStatus(resp *http.Response, err error, site string) string {
	if resp != nil {
		if resp.StatusCode >= 200 && resp.StatusCode <= 300 {
			registraLog(site, true)
			return "O site " + site + " foi carregado com sucesso"

		} else if resp.StatusCode > 300 && resp.StatusCode < 600 {
			registraLog(site, false)
			return "O site " + site + " não foi carregado com sucesso com o status " + strconv.Itoa(resp.StatusCode)
		} else {
			registraLog(site, false)
			return errorResponse(site)
		}
	} else {
		registraLog(site, false)
		return errorResponse(site)
	}

}
func errorResponse(site string) string {
	return "O site " + site + " não foi possível encontrar o site"
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	} else {
		arquivo.WriteString(time.Now().Format("01-02-2006 15:04:05") + " o  site: " + strings.Replace(site, "\n", "", -1) + " - online: " + strconv.FormatBool(status) + "\n")
	}

	fmt.Println(arquivo)
	arquivo.Close()
}
func mostrarLogs() {
	arquivo, err := os.Open("log.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		leitor := bufio.NewReader(arquivo)
		for {
			linha, err2 := leitor.ReadString('\n')
			if err2 == nil {
				fmt.Print(linha)
			} else {
				break
			}
		}
	}

}
