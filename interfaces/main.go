package main

import (
	"fmt"

	"github.com/humbertodias/go-course/interfaces/model"
)

func main() {
	fmt.Println("OK")

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

	queroAcordarComUmCacarejo(jojo)
	queroOuvirUmaPataNoLago(jojo)
}

func queroAcordarComUmCacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmaPataNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
