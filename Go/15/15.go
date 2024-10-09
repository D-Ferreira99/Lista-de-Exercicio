package main

import (
	"fmt"
)

type SaldoInsuficienteError struct {
	saldo  float64
	valor  float64
}

func (e *SaldoInsuficienteError) Error() string {
	return fmt.Sprintf("saldo insuficiente: disponível R$%.2f\nSaque solicitado R$%.2f", e.saldo ,e.valor)
}

type ContaBancaria struct {
	titular string
	saldo   float64
}

func (c *ContaBancaria) Depositar(valor float64) {
	c.saldo += valor
}

func (c *ContaBancaria) Sacar(valor float64) error {
	if valor > c.saldo {
		return &SaldoInsuficienteError{saldo: c.saldo, valor: valor}
	}
	c.saldo -= valor
	return nil
}

func main() {

	conta := ContaBancaria{titular: "Thomas", saldo: 20.00}

	if err := conta.Sacar(50.00); err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Saque realizado com sucesso!")
	}

	conta.Depositar(100.00)

	if err := conta.Sacar(80.00); err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Saque realizado com sucesso!")
	}
}
