// Package service contém as regras de serviço da entidade item
package service

import (
	"desafiotecnico/model"
	"desafiotecnico/util"
	"fmt"
)

// Separar Valores Por Email
// Uma lista de compras e uma lista de emails são passadas para serem tratadas
// Se ela seguir as regras definidas, é retornado a lista com valores separados por email e false
// Se ela não seguir as regras definidas, é retornado nil e true
func SepararValoresPorEmail(listaDeCompras []model.Item, listaDeEmails []string) (map[string]int, util.ErrorType) {

	err := verificarTamanhoDasListas(listaDeCompras, listaDeEmails)

	if err.Status {
		return nil, err
	}

	valorTotal, err := somarValores(listaDeCompras)

	if err.Status {
		return nil, err
	}

	listaDeValores, err := dividirValores(valorTotal, listaDeEmails)

	if err.Status {
		return nil, err
	}

	listaEmailEValores, err := mapearListaDeEmailsEValores(listaDeValores, listaDeEmails)

	if err.Status {
		return nil, err
	}

	return listaEmailEValores, util.ErrorType{Status: false}
}

// Verificar tamanho das listas
// Recebe uma lista de compras e uma lista de emails e verifica se os tamanhos são válidos
// Se a operação for válida, retorna false
// Se a operação for inválida, retorna true
func verificarTamanhoDasListas(listaDeCompras []model.Item, listaDeEmails []string) util.ErrorType {

	if len(listaDeCompras) > 0 && len(listaDeEmails) > 0 {
		return util.ErrorType{Status: false}
	}

	switch {
	case len(listaDeCompras) <= 0 && len(listaDeEmails) <= 0:
		return util.ErrorType{Mensagem: "Lista de compras e emails inválida!", Status: true}
	case len(listaDeCompras) <= 0:
		return util.ErrorType{Mensagem: "Lista de compras inválida!", Status: true}
	default:
		return util.ErrorType{Mensagem: "Lista de emails inválida!", Status: true}
	}
}

// Somar Valores
// Recebe uma lista de compras e calcula a soma dos itens dessa lista
// Se a operação for válida, retorna o valor total e false
// Se a operação for inválida, retorna 0 e true
func somarValores(listaDeCompras []model.Item) (int, util.ErrorType) {

	var valorTotal int

	for _, item := range listaDeCompras {
		valorTotal += item.Preco * item.Quantidade
	}

	if valorTotal <= 0 {
		return 0, util.ErrorType{Mensagem: "Valor total inválido!", Status: true}
	}

	fmt.Println("Valor total dos itens:", valorTotal)

	return valorTotal, util.ErrorType{Status: false}
}

// Dividir Valores
// Recebe o valor total dos itens e uma lista de emails
// Se a operação for válida, retorna uma lista de valores que representa a divisão do valor total
// pela quantidade emails e false
// Se a operação for inválida, retorna nil e true
func dividirValores(valorTotal int, listaDeEmails []string) ([]int, util.ErrorType) {

	switch {
	case len(listaDeEmails) <= 0:
		return nil, util.ErrorType{Mensagem: "Lista de emails vazia!", Status: true}
	case valorTotal < len(listaDeEmails):
		return nil, util.ErrorType{Mensagem: "Valor total menor que a quantidade de emails!", Status: true}
	}

	listaDeValores, err := dividirPorModulo0(valorTotal, listaDeEmails)

	if err.Status == false {
		return listaDeValores, err
	}

	listaDeValores, err = dividirPorModuloDiferenteDe0(valorTotal, listaDeEmails)

	return listaDeValores, err
}

// Dividir Por Modulo 0
// Recebe o valor total e uma lista de emails, verifica se a divisão é modulo 0 e retorna uma lista de valores
// Se a operação for válida, retorna uma lista de valores e false
// Se a operação não for válida, retorna nil e true
func dividirPorModulo0(valorTotal int, listaDeEmails []string) ([]int, util.ErrorType) {

	if valorTotal%len(listaDeEmails) == 0 {

		valorPorUnidade := valorTotal / len(listaDeEmails)

		var listaDeValores []int

		for i := 0; i < len(listaDeEmails); i++ {
			listaDeValores = append(listaDeValores, valorPorUnidade)
		}

		return listaDeValores, util.ErrorType{Status: false}
	}

	return nil, util.ErrorType{Status: true}
}

// Dividir Por Modulo Diferente de 0
// Recebe o valor total e uma lista de emails, verifica se a divisão é modulo diferente de 0 e retorna uma lista de valores
func dividirPorModuloDiferenteDe0(valorTotal int, listaDeEmails []string) ([]int, util.ErrorType) {

	modulo := valorTotal % len(listaDeEmails)
	valorPorUnidade := (valorTotal - modulo) / len(listaDeEmails)

	var listaDeValores []int

	for i := 0; i < len(listaDeEmails); i++ {
		if modulo > 0 {
			listaDeValores = append(listaDeValores, (valorPorUnidade + (modulo - (modulo - 1))))
			modulo -= 1
		} else {
			listaDeValores = append(listaDeValores, valorPorUnidade)
		}
	}

	return listaDeValores, util.ErrorType{Status: false}
}

// Mapear Lista De Emails E Valores
// Recebe uma lista de valores e uma lista de emails
// Se a operação for válida, retorna um map com chave (email) e valor (valor por unidade) e false
// Se a operação não for válida, retorna nil e true
func mapearListaDeEmailsEValores(listaDeValores []int, listaDeEmails []string) (map[string]int, util.ErrorType) {

	if len(listaDeValores) != len(listaDeEmails) {
		return nil, util.ErrorType{Mensagem: "A lista de valores e a lista de emails não possuem o mesmo tamanho!", Status: true}
	}

	listaEmailsEValores := make(map[string]int)

	for i, valor := range listaDeValores {
		listaEmailsEValores[listaDeEmails[i]] = valor
	}

	return listaEmailsEValores, util.ErrorType{Status: false}
}
