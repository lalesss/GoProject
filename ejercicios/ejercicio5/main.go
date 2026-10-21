package main

import (
	"fmt"
	"sync"
)

type Servidor struct {
	Nombre, IP string
	Activo     bool
}

func main() {
	canal := make(chan *Servidor, 2)
	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s := &Servidor{fmt.Sprint("Servidor ", i),
				fmt.Sprint("192.168.1.", 10+i), true}
			canal <- s
		}(i)
	}

	wg.Wait()
	close(canal)
	for s := range canal {
		fmt.Println(*s)
	}
}
