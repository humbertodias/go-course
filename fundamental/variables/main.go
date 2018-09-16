package main

import (
	"fmt"
)

// group of variables
var (
	Address string        // public
	phone   = "9999-9999" // private
	amount  int           // = 0
	bought  bool          // false
	value   float64       // 0.00
	words   rune
)

func main() {
	test := 9 // Declaration and attribution
	fmt.Printf("Address: %s\r\n", Address)
	fmt.Printf("Amount: %d\r\n", amount)
	fmt.Printf("Bought: %v\r\n", bought)
	fmt.Printf("Words: %v\r\n", words)
	fmt.Printf("Test: %d\r\n", test)
}
