package main

import "fmt"

type Imprimivel interface {
	Imprimir()
}

type Relatorio struct {
	conteudo string
}

func (r Relatorio) Imprimir() {
	fmt.Println("Imprimindo Relatório:", r.conteudo)
}

type Contrato struct {
	detalhes string
}

func (c Contrato) Imprimir() {
	fmt.Println("Imprimindo Contrato:", c.detalhes)
}

func main() {
	var itens []Imprimivel
	itens = append(itens, Relatorio{conteudo: "Relatório de Vendas"})
	itens = append(itens, Contrato{detalhes: "Contrato de Trabalho"})

	for _, item := range itens {
		item.Imprimir()
	}
}
