package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/humbertodias/go-course/web-post/model"
)

// https://requestbin.fullcontact.com/
func main() {

	// user
	user := model.User{}
	user.ID = 1
	user.Name = "Humberto"

	sendContent, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Erro no marshal do usuario. Error: ", err.Error())
		return
	}

	request, err := http.NewRequest("POST", "http://requestbin.fullcontact.com/116vr5r1", bytes.NewBuffer(sendContent))
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
		fmt.Println(string(body))
	}

}
