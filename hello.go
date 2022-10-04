package main

import "fmt"


func main()	{
	 nome := "JP"
	 versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("0- Sair do programa")
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")


	var command int
	fmt.Scan(&command)

	fmt.Println("O endereço da variável é", &command)
	fmt.Println("o comando escolhido foi", command)
}

