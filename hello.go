package main

import "fmt"

func main() {
	nome := "JP"
	versao := 1.1

	fmt.Println("Hi, sr.", nome)
	fmt.Println("Version", versao)

	fmt.Println("0- Exit program")
	fmt.Println("1- Start monitoring")
	fmt.Println("2- Show logs")

	var command int
	fmt.Scan(&command)

	fmt.Println("O endereço da variável é", &command)
	fmt.Println("o comando escolhido foi", command)

	// if command == 1 {
	// 	fmt.Println("Monitorando")
	// } else if command == 2 {
	// 	fmt.Println("Show Logs")
	// } else if command == 3 {
	// 	println("Exit")
	// } else {
	// 	fmt.Println("Command not found")
	// }

	switch command {
	case 1:
		fmt.Println("Monitoring")
	case 2:
		fmt.Println("Show logs")
	case 3:
		fmt.Println("Exit")
	default:
		fmt.Println("Command not found")
	}
}
