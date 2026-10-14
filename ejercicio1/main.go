package main

import "fmt"

func main() {
	numero := 42
	ptr := &numero

	fmt.Println(numero, ptr, *ptr)

	*ptr = 100
	fmt.Println(numero)
}
