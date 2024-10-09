package main

import (
	"errors"
	"fmt"
)

type Conta struct {
	titular string
	saldo   float64 
}

func (c *Conta) Depositar(valor float64) {
	if valor > 0 {
		c.saldo += valor
	} else {
		fmt.Println("O valor do depósito deve ser um valor positivo.")
	}
}

func (c *Conta) Sacar(valor float64) error {
	if valor > 0 {
		if valor <= c.saldo {
			c.saldo -= valor
			return nil
		} else {
			return errors.New("saldo insuficiente")
		}
	}
	return errors.New("o valor do saque deve ser um valor positivo")
}


func (c *Conta) ExibirSaldo() float64 {
	return c.saldo
}

func (c *Conta) DefinirTitular(titular string) {
	c.titular = titular
}

func (c *Conta) ObterTitular() string {
	return c.titular
}

func main() {

	conta := Conta{}
	conta.DefinirTitular("Roberto")

	fmt.Println("Titular:", conta.ObterTitular())

	conta.Depositar(2000.00)
	fmt.Printf("Saldo após depósito: %.2f\n", conta.ExibirSaldo())

	err := conta.Sacar(500.00)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Saldo após saque: %.2f\n", conta.ExibirSaldo())
	}

	err = conta.Sacar(2000.00)
	if err != nil {
		fmt.Println(err)
	}
}
