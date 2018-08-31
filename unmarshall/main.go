package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/humbertodias/go-course/unmarshall/model"
)

// https://jsonplaceholder.typicode.com/
func main() {

	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("Erro ao abrir pagina do google. Error: ", err.Error())
		return
	}
	request.SetBasicAuth("fizz", "buzz")
	request.Header.Set("content-type", "application/json; charset=utf-8")

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Erro na responsta do POST: ", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Erro ao ler o conteudo resposta: ", err.Error())
			return
		}
		// fmt.Println(string(body))
		post := model.BlogPost{}
		err = json.Unmarshal(body, &post)
		if err != nil {
			fmt.Println("Erro ao converter o retorno do servidor: ", err.Error())
			return
		}

		fmt.Printf("Post é: %+v\r\n", post)
	}

}
