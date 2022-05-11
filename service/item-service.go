// Package service contém as regras de serviço da entidade item
package service

import (
	"desafiotecnico/model"
	"fmt"
)

// Separar Valores Por Email
// Uma lista de compras e uma lista de emails são passadas para serem tratadas
// Se ela seguir as regras definidas, é retornado a lista com valores separados por email e true
// Se ela não seguir as regras definidas, é retornado nil e false
func SepararValoresPorEmail(listaDeCompras []model.Item, listaDeEmails []string) (map[string]int, bool) {

	if len(listaDeCompras) > 0 && len(listaDeEmails) > 0 {

		valorTotal, isValorTotalValido := somaDosValores(listaDeCompras)

		if isValorTotalValido {

			listaDeValores, isListaDeValoresValida := dividirValores(valorTotal, listaDeEmails)

			if isListaDeValoresValida {

				listaEmailEValores, isListaEmailEValoresValida := listaDeEmailsEValores(listaDeValores, listaDeEmails)

				if isListaEmailEValoresValida {
					fmt.Println("Valor total dos itens:", valorTotal)
					return listaEmailEValores, true
				} else {
					fmt.Println("Lista de email e valores inválida!")
					return nil, false
				}

			} else {
				fmt.Println("Lista de valores inválida!", listaDeValores)
				return nil, false
			}

		} else {
			fmt.Println("Valor total inválido!", valorTotal)
			return nil, false
		}

	} else {
		if len(listaDeCompras) <= 0 && len(listaDeEmails) <= 0 {
			fmt.Println("Lista de compras e emails inválida!", listaDeCompras, listaDeEmails)
			return nil, false
		} else if len(listaDeCompras) <= 0 {
			fmt.Println("Lista de compras inválida!", listaDeCompras)
			return nil, false
		} else {
			fmt.Println("Lista de emails inválida!", listaDeEmails)
			return nil, false
		}
	}
}

// Soma Dos Valores
// Recebe uma lista de compras e calcula a soma dos itens dessa lista
// Se a operação for válida, retorna o valor total e true
// Se a operação for inválida, retorna 0 e false
func somaDosValores(listaDeCompras []model.Item) (int, bool) {

	var valorTotal int

	for _, item := range listaDeCompras {
		valorTotal += item.Preco * item.Quantidade
	}

	if valorTotal <= 0 {
		return 0, false
	}

	return valorTotal, true
}

// Dividir Valores
// Recebe o valor total dos itens e uma lista de emails
// Se a operação for válida, retorna uma lista de valores que representa a divisão do valor total
// pela quantidade emails e true
// Se a operação for inválida, retorna nil e false
func dividirValores(valorTotal int, listaDeEmails []string) ([]int, bool) {

	var listaDeValores []int

	if len(listaDeEmails) > 0 && valorTotal >= len(listaDeEmails) {

		// Verifica se é módulo 0
		if valorTotal%len(listaDeEmails) == 0 {

			valorPorUnidade := valorTotal / len(listaDeEmails)

			for i := 0; i < len(listaDeEmails); i++ {
				listaDeValores = append(listaDeValores, valorPorUnidade)
			}
		} else {
			modulo := valorTotal % len(listaDeEmails)
			valorPorUnidade := (valorTotal - modulo) / len(listaDeEmails)

			for i := 0; i < len(listaDeEmails); i++ {
				if modulo > 0 {
					listaDeValores = append(listaDeValores, (valorPorUnidade + (modulo - (modulo - 1))))
					modulo -= 1
				} else {
					listaDeValores = append(listaDeValores, valorPorUnidade)
				}
			}
		}
	} else {
		if len(listaDeEmails) <= 0 {
			fmt.Println("Lista de emails vazia!")
		}

		if valorTotal < len(listaDeEmails) {
			fmt.Println("Valor total menor que a quantidade de emails!")
		}

		return nil, false
	}

	return listaDeValores, true
}

// Lista De Emails E Valores
// Recebe uma lista de valores e uma lista de emails
// Se a operação for válida, retorna um map com chave (email) e valor (valor por unidade) e true
// Se a operação não for válida, nil e false
func listaDeEmailsEValores(listaDeValores []int, listaDeEmails []string) (map[string]int, bool) {

	listaEmailsEValores := make(map[string]int)

	if len(listaDeValores) == len(listaDeEmails) {
		for i, valor := range listaDeValores {
			listaEmailsEValores[listaDeEmails[i]] = valor
		}
	} else {
		fmt.Println("A lista de valores e a lista de emails não possuem o mesmo tamanho!")
		return nil, false
	}

	return listaEmailsEValores, true
}
