package main

import (
	"fmt"
	"sync"
)

type Banco struct {
	Nome string
}

var instance *Banco

var once sync.Once

func GetBanco() *Banco {
	once.Do(func() {
		instance = &Banco{Nome: "Banco de Dados Exemplo"}
	})
	return instance
}

func main() {
	banco1 := GetBanco()
	fmt.Println("Banco A:", banco1.Nome)

	banco2 := GetBanco()
	fmt.Println("Banco B:", banco2.Nome)

	fmt.Println("As instâncias são iguais:", banco1 == banco2)
}
