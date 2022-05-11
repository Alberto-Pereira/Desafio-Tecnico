// Package model contém item, que representa a entidade principal do sistema
package model

// Item é uma entidade usada em uma lista de compras
type Item struct {
	Nome       string // Nome do item
	Quantidade int    // Quantidade do item
	Preco      int    // Preco em centavos do item
}
