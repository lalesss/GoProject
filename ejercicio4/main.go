package main

import "fmt"

func duplicar(ptr *int) {
	if ptr == nil {
		fmt.Println("Error: puntero nulo")
		return
	}
	*ptr *= 2
}

func main() {
	numero := 10

	duplicar(&numero)
	fmt.Println("Valor duplicado:", numero)

	duplicar(nil)
}
