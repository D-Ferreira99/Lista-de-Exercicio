package main

import (
	"fmt"
)

type Calculo struct{}

func (m Calculo) Fatorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * m.Fatorial(n-1)
}
func main() {
	mat := Calculo{}

	num := 8
	resultado := mat.Fatorial(num)
	fmt.Printf("O fatorial de %d: %d\n", num, resultado)
}
