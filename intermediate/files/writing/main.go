package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/humbertodias/go-course/intermediate/files/model"
)

func main() {

	fullpath := "intermediate/files/reading/cidades.csv"
	arquivo, err := os.Open(fullpath)
	if err != nil {
		fmt.Println("[main] Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}
	defer arquivo.Close()

	leitorCsv := csv.NewReader(arquivo)
	conteudo, err := leitorCsv.ReadAll()
	if err != nil {
		fmt.Println("[main] Houve um erro ao ler o arquivo com leitor CSV. Erro: ", err.Error())
		return
	}

	fullJSONPath := "intermediate/files/reading/cidades.json"
	arquivoJSON, err := os.Create(fullJSONPath)
	if err != nil {
		fmt.Println("[main] Houve um erro ao criar o arquivo JSON. Erro: ", err.Error())
		return
	}
	defer arquivoJSON.Close()

	escritor := bufio.NewWriter(arquivoJSON)
	escritor.WriteString("[\r\n")
	for _, linha := range conteudo {
		for indiceItem, item := range linha {
			dados := strings.Split(item, "/")
			cidade := model.City{}
			cidade.Name = dados[0]
			cidade.State = dados[1]
			fmt.Printf("Cidade: %+v\r\n", cidade)
			cidadeJSON, err := json.Marshal(cidade)
			if err != nil {
				fmt.Println("[main] Houve um erro ao gerar o json do item ", item, ". Erro: ", err.Error())
			}
			escritor.WriteString("  " + string(cidadeJSON))
			if (indiceItem + 1) < len(linha) {
				escritor.WriteString(",\r\n")
			}
		}
	}
	escritor.WriteString("\r\n]")
	escritor.Flush()
}
