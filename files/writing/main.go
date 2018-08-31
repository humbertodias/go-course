package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/humbertodias/go-course/files/model"
)

func main() {
	jsonFile, err := os.Create("city.json")
	if err != nil {
		fmt.Println("Erro criando arquivo json", err.Error())
		return
	}
	// Marca para fechar arquivo ao final da execucao
	defer jsonFile.Close()

	writer := bufio.NewWriter(jsonFile)
	writer.WriteString("[\r\n")

	cities := []string{"Sao Paulo", "Rio de Janeiro"}

	for _, line := range cities {
		city := model.City{}
		city.Name = line

		fmt.Println(city)

		cityJSON, err := json.Marshal(city)
		if err != nil {
			fmt.Println("Erro criando arquivo json", err.Error())
		}
		writer.WriteString(" " + string(cityJSON))
	}
	writer.WriteString("\r\n]")
	writer.Flush()

}
