package main

import "fmt"

type CuentaBancaria struct {
	saldo float64
}

func (c *CuentaBancaria) Depositar(monto float64) {
	if monto > 0 {
		c.saldo += monto
	}
}

func (c *CuentaBancaria) Retirar(monto float64) {
	if monto <= 0 || monto > c.saldo {
		fmt.Println("Error al retirar")
		return
	}
	c.saldo -= monto
}

func main() {
	cuenta := CuentaBancaria{1000}

	cuenta.Depositar(500)
	cuenta.Retirar(200)

	fmt.Println("Saldo final:", cuenta.saldo)
}
