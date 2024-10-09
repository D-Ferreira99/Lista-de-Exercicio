package main

import "fmt"

type Animal interface {
	Som() string
}

type Cachorro struct{}

func (c Cachorro) Som() string {
	return "Au Au"
}

type Gato struct{}

func (g Gato) Som() string {
	return "Miau"
}

func somAnimal(animal Animal) {
	fmt.Println(animal.Som())
}

func main() {

	cachorro := Cachorro{}
	gato := Gato{}

	somAnimal(cachorro)
	somAnimal(gato)
}
