package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	response, err := client.Get("https://www.google.es")
	if err != nil {
		fmt.Println("Erro ao abrir pagina do google. Error: ", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Erro na responsta: ", err.Error())
			return
		}
		fmt.Println(string(body))
	}

	request, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		fmt.Println("Erro na responsta: ", err.Error())
		return
	}
	request.SetBasicAuth("test", "test")
	response, err = client.Do(request)

	if err != nil {
		fmt.Println("Erro na responsta: ", err.Error())
		return
	}

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Erro na resposta: ", err.Error())
			return
		}
		fmt.Println("SECOND REQUEST")
		fmt.Println(string(body))
	}

}
