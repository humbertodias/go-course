package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("cidades.csv")

	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro:", err.Error())
		return
	}
	defer file.Close()
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	fmt.Println("O conteudo da linha é: ", line)
	// }

	csvReader := csv.NewReader(file)
	content, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro:", err.Error())
		return
	}

	for indexLine, line := range content {
		fmt.Printf("Line [%d] is %s\r\n", indexLine, line)
		for indexItem, item := range line {
			fmt.Printf("Item [%d] is %s\r\n", indexItem, item)
		}
	}
}
