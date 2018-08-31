package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/humbertodias/go-course/web-server/model"
)

// Hello model
func Hello(w http.ResponseWriter, r *http.Request) {
	time := time.Now().Format("15:04:05")
	page := model.Page{}
	page.Time = time
	if err := Models.ExecuteTemplate(w, "hello.html", page); err != nil {
		http.Error(w, "Houve um erro na renderizacao da pagina", http.StatusInternalServerError)
		fmt.Println("Hello erro na execucao do modelo")
	}
}
