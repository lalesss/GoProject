package main

import "fmt"

type Libro struct {
	titulo string
	autor  string
	precio float64
}

func aplicarDescuento(libro *Libro) {
	libro.precio *= 0.8
}

func main() {
	libro := Libro{"El principito", "Antoine de Saint-Exupery", 200}

	fmt.Println("Precio inicial:", libro.precio)
	aplicarDescuento(&libro)
	fmt.Println("Precio final:", libro.precio)
}
