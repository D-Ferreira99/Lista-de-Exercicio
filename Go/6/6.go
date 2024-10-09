package main

import "fmt"

type Motor struct {
	potencia int 
	tipo     string 
}

func (m Motor) Exibir() string {
	return fmt.Sprintf("Motor: %s, Potência: %d CV", m.tipo, m.potencia)
}

type Carro struct {
	marca string
	modelo string
	ano int
	motor Motor 
}

func (c Carro) Exibir() string {
	return fmt.Sprintf("Carro: %s %s (%d): %s", c.marca, c.modelo, c.ano, c.motor.Exibir())
}

func main() {
	motor1 := Motor{potencia: 86, tipo: "Gasolina"}

	carro1 := Carro{marca: "Fiat", modelo: "Strada", ano: 2015, motor: motor1}

	fmt.Println(carro1.Exibir())
}
