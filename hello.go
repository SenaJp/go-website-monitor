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

const monitoring = 3
const delay = 5

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
			showLogs()
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
	fmt.Println("")

	return command
}

func startMonitoring() {
	fmt.Println("Monitoring...")
	sites := readSitesToArchive()


	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Checking position", i, "the site is:", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testSite(site string) {
	resp, err := http.Get(site)

	if(err != nil) {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso")
		logRegister(site, true)
	} else {
		fmt.Println("O site", site, "está com algum problema. StatusCode:", resp.StatusCode)
		logRegister(site, false)
	}
}

func readSitesToArchive() []string {
	var sites []string

	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)
	
		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func logRegister(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("O erro encontrado foi:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func showLogs(){
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("O erro encontrado foi:", err)
	}

	fmt.Println(string(file))
}