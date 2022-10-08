package main

import "fmt"
import "os"

func main() {

	showIntroduction()
	showMenu()
	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoring")
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
	fmt.Println("o comando escolhido foi", command)

	return command
}


