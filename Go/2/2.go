package main

import "fmt"

type Carro struct {
	velocidade int
}

func (c *Carro) Aceleracao(aumento int) {
	c.velocidade += aumento
}

func (c *Carro) Frenagem(reducao int) {
	c.velocidade -= reducao
	if c.velocidade < 0 {
		c.velocidade = 0
	}
}

func (c *Carro) ExibirVelocidade() {
	fmt.Printf("A velocidade atual é: %d km/h\n", c.velocidade)
}

func main() {

	carro := Carro{velocidade: 0}

	carro.Aceleracao(80)
	carro.ExibirVelocidade()

	carro.Frenagem(20)
	carro.ExibirVelocidade()

	carro.Frenagem(15)
	carro.ExibirVelocidade()
}
