package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	showIntroduction()
	for {
		showMenu()
		command := readCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Show logs")
		case 0:
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Command not found")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	nome := "JP"
	versao := 1.1
	fmt.Println("Hi, sr.", nome)
	fmt.Println("Version", versao)
}

func showMenu() {
	fmt.Println("0- Exit program")
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	site := "https://random-status-code.herokuapp.com"
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso")
	} else {
		fmt.Println("O site", site, "está com algum problema. StatusCode:", resp.StatusCode)
	}
}
