package main

import (
	"fmt"
	"os"

	"github.com/tourist/utils"
)

type Routes struct {
	origem  string
	destino string
	valor   int
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Favor mencionar o arquivo de CSV \n Exemplo: go run main.go input-routes.csv")
	} else {

		fileName := os.Args[1]

		lines, err := utils.ReadCsv(fileName)

		if err != nil {
			fmt.Printf("Arquivo %s inválido ! \n Favor escolher um arquivo csv valido.", fileName)
		}

		fmt.Println(ShowWelcome())

		fmt.Println(lines)
	}
}

func ShowWelcome() string {
	return "Bem vindo ao programa para turistas"
}
