package main

import "fmt"

// Definimos el struct CuentaBancaria
type CuentaBancaria struct {
	Titular string
	Saldo   float64
}

// Método para depositar dinero (receptor puntero para modificar saldo)
func (c *CuentaBancaria) Depositar(monto float64) {
	c.Saldo += monto
	fmt.Printf("Depositados %.2f. Nuevo saldo: %.2f\n", monto, c.Saldo)
}

// Método para retirar dinero (receptor puntero)
func (c *CuentaBancaria) Retirar(monto float64) {
	if monto > c.Saldo {
		fmt.Println("Fondos insuficientes para retirar.")
		return
	}
	c.Saldo -= monto
	fmt.Printf("Retirados %.2f. Nuevo saldo: %.2f\n", monto, c.Saldo)
}

// Struct para CuentaDeAhorros
type CuentaDeAhorros struct {
	Titular     string
	Saldo       float64
	TasaInteres float64
}

func (c *CuentaDeAhorros) Depositar(monto float64) {
	c.Saldo += monto
	fmt.Printf("Depositados %.2f a CuentaDeAhorros. Nuevo saldo: %.2f\n", monto, c.Saldo)
}

func (c *CuentaDeAhorros) Retirar(monto float64) {
	if monto > c.Saldo {
		fmt.Println("Fondos insuficientes para retirar en CuentaDeAhorros.")
		return
	}
	c.Saldo -= monto
	fmt.Printf("Retirados %.2f de CuentaDeAhorros. Nuevo saldo: %.2f\n", monto, c.Saldo)
}

// Interfaz Operable con métodos para depositar y retirar
type Operable interface {
	Depositar(monto float64)
	Retirar(monto float64)
}

// Función que recibe cualquier Operable y hace operaciones
func RealizarOperaciones(op Operable) {
	op.Depositar(100)
	op.Retirar(50)
}

func main() {
	cuenta := CuentaBancaria{Titular: "Ana", Saldo: 1000}
	ahorro := CuentaDeAhorros{Titular: "Luis", Saldo: 5000, TasaInteres: 0.02}

	fmt.Println("Operando sobre CuentaBancaria:")
	RealizarOperaciones(&cuenta)

	fmt.Println("\nOperando sobre CuentaDeAhorros:")
	RealizarOperaciones(&ahorro)
}
