package main

import "fmt"

func main() {
	a := 5
	b := 10

	temporal := a
	a = b
	b = temporal

	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
