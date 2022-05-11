// Package util contém arquivos que auxiliam na criação de casos/testes para a aplicação.
package util

import (
	"desafiotecnico/model"
	"fmt"
	"math/rand"
	"time"

	"github.com/thanhpk/randstr"
)

// Gerar Itens
// Gera uma quantidade randômica de itens e retorna em uma lista
// É possível alterar algumas propriedades
func GerarItens() []model.Item {
	var listaDeCompras []model.Item

	var quantidadeDeItensASerGerada int = 100       // Escolha até quantos itens serão gerados na lista de compras | 0 até n
	var quantidadeDeCaracteresDoNomeDoItem int = 10 // Escolha a quantidade de caracteres do item | Valor n será fixo

	seed := rand.NewSource(time.Now().UnixNano() * time.Now().UnixNano()) // Gera uma seed
	random := rand.New(seed)                                              // Atribue a variável random um Rand usando a seed passada
	quantidadeDeItens := random.Intn(quantidadeDeItensASerGerada)         // Gera uma quantidade de itens randômica baseada na variável random

	var quantidadeMaxDeCadaItem int = 10 // Escolha qual será a quantidade máxima de cada item
	var precoMaxDeCadaItem int = 1000    // Escolha qual será o preço por unidade máximo de cada item | Valor em centavos

	// Cria uma lista de compras com valores gerados randomicamente
	for i := 0; i < quantidadeDeItens; i++ {

		quantidadeDoItemAtual := random.Intn(quantidadeMaxDeCadaItem)
		precoDoItemAtual := random.Intn(precoMaxDeCadaItem)

		listaDeCompras = append(listaDeCompras, model.Item{
			Nome:       randstr.String(quantidadeDeCaracteresDoNomeDoItem),
			Quantidade: quantidadeDoItemAtual,
			Preco:      precoDoItemAtual,
		})

	}

	if len(listaDeCompras) <= 0 {
		fmt.Println("Lista de compras gerada se encontra vazia!")
		return nil
	} else {
		fmt.Println("Quantidade de itens gerados:", quantidadeDeItens)
	}

	return listaDeCompras
}
