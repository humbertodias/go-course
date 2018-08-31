package main

import (
	"fmt"
)

func main() {
	var test [3]int
	test[0] = 0
	test[1] = 1
	test[2] = 2
	fmt.Printf("len(array)=%d\r\n", len(test))
	fmt.Printf("Array[0]=%d\r\n", test[0])

	cantores := []string{"um", "dois"}

	fmt.Printf("cantores %v\r\n", cantores)

	capitais := []string{"Lisboa", "Brasilia", "Luanda"}

	for indice, cidade := range capitais {
		fmt.Printf("%d %v\r\n", indice, cidade)
	}
}
