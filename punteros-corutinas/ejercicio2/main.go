package main

import (
	"fmt"
	"sync"
)

type Resultado struct {
	Suma       int
	Completado bool
}

func main() {
	res := Resultado{}
	numeros := []int{10, 20, 30, 40, 50}
	var wg sync.WaitGroup

	wg.Add(1)
	go func(p *Resultado) {
		defer wg.Done()
		for _, n := range numeros {
			p.Suma += n
		}
		p.Completado = true
	}(&res)

	wg.Wait()
	fmt.Println(res)
}
