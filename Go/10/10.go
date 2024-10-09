package main

import "fmt"

type Calcular struct{}

func (c Calcular) SomarDois(num1, num2 float64) float64 {
	return num1 + num2
}

func (c Calcular) SomarTres(num1, num2, num3 float64) float64 {
	return num1 + num2 + num3
}

func main() {
	calculadora := Calcular{}


	resultadoDois := calculadora.SomarDois(45, 580)
	fmt.Printf("Resultado da soma com dois numeros: %.2f\n", resultadoDois)

	resultadoTres := calculadora.SomarTres(70, 40, 220)
	fmt.Printf("Resultado da soma com três números: %.2f\n", resultadoTres)
}
