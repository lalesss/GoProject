package main

import (
	"fmt"
	"sync"
)

func main() {
	datos := make([]int, 5)
	var wg sync.WaitGroup

	for i := range datos {
		wg.Add(1)
		go func(i int, p *int) {
			defer wg.Done()
			*p = i * 10
		}(i, &datos[i])
	}

	wg.Wait()
	fmt.Println(datos)
}
