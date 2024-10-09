package main

import "fmt"

type Carro struct {
    marca  string
    modelo string
    ano    int
}
func main() {
    veiculo1 := Carro{"Lamborghini", "Temerario", 2024}
    veiculo2 := Carro{"Fiat", "Strada", 2015}
    veiculo3 := Carro{"Chevrolet", "Onix", 2019}
   
    fmt.Println(veiculo1)
    fmt.Println(veiculo2)
    fmt.Println(veiculo3)
}