package main

import (
	"fmt"
)

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums))

	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))

	capitais = append(capitais, "Brasilia")
	fmt.Println(capitais, len(capitais), cap(capitais))
	capitais[0] = "Nova York"

	for indice, capital := range capitais {
		fmt.Printf("%d %s\r\n", indice, capital)
	}

	cidades := make([]string, 5)

	cidades[0] = "Nova York"
	cidades[1] = "Londres"
	cidades[2] = "Madri"
	cidades[3] = "Tokio"
	cidades[4] = "Singapura"
	for indice, cidade := range cidades {
		fmt.Printf("Cidade[%d] %s\r\n", indice, cidade)
	}
	cidadesAsia := cidades[2:]
	for indice, cidade := range cidadesAsia {
		fmt.Printf("CidadeAsia[%d] %s\r\n", indice, cidade)
	}

}
