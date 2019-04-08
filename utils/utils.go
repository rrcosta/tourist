package utils

import (
	"encoding/csv"
	"log"
	"os"
)

//  ReadCsv is a function that open the file
func ReadCsv(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal("Erro ao abrir o arquivo de entrada \n Detalhes: ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.TrimLeadingSpace = true

	lines, err := reader.ReadAll()

	if err != nil {
		log.Fatal("Error reading all lines: ", err)
	}

	return lines, err
}
