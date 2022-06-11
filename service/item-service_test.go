// Package service contém as regras de serviço da entidade item
package service

import (
	"desafiotecnico/model"
	"desafiotecnico/util"
	"testing"

	"github.com/stretchr/testify/require"
)

// Separar Valores Por Email
// Lista De Compras E Lista De Emails Valida - Modulo 0
// Deve fazer a separação de valores por email possuindo um módulo 0
func TestSepararValoresPorEmail_ListaDeComprasEListaDeEmailsValida_Modulo0(t *testing.T) {

	assertions := require.New(t)

	// Módulo 0
	listaDeCompras := []model.Item{
		{Nome: "Detergente", Quantidade: 4, Preco: 500},
		{Nome: "Uva", Quantidade: 6, Preco: 750},
		{Nome: "Arroz", Quantidade: 2, Preco: 1000},
	}

	listaDeEmails := []string{
		"user1@gmail.com", "user2@gmail.com", "user3@gmail.com", "user4@gmail.com", "user5@gmail.com",
	}

	listaEmailEValores, isListaEmailEValoresValida := SepararValoresPorEmail(listaDeCompras, listaDeEmails)

	assertions.NotEmpty(listaEmailEValores)
	assertions.Equal(true, isListaEmailEValoresValida)

}

// Separar Valores Por Email
// Lista De Compras E Lista De Emails Valida - Modulo Diferente De 0
// Deve fazer a separação de valores por email possuindo um módulo diferente 0
func TestSepararValoresPorEmail_ListaDeComprasEListaDeEmailsValida_ModuloDiferenteDe0(t *testing.T) {

	assertions := require.New(t)

	// Módulo diferente de 0
	listaDeCompras := []model.Item{
		{Nome: "Detergente", Quantidade: 2, Preco: 500},
		{Nome: "Uva", Quantidade: 5, Preco: 750},
		{Nome: "Arroz", Quantidade: 1, Preco: 1000},
	}

	listaDeEmails := []string{
		"user1@gmail.com", "user2@gmail.com", "user3@gmail.com", "user4@gmail.com", "user5@gmail.com", "user6@gmail.com",
	}

	listaEmailEValores, isListaEmailEValoresValida := SepararValoresPorEmail(listaDeCompras, listaDeEmails)

	assertions.NotEmpty(listaEmailEValores)
	assertions.Equal(true, isListaEmailEValoresValida)

}

// Separar Valores Por Email
// Valor Total Invalido
// Deve retornar uma lista de valores vazia
func TestSepararValoresPorEmail_ValorTotalInvalido(t *testing.T) {

	assertions := require.New(t)

	listaDeCompras := []model.Item{
		{Nome: "Detergente", Quantidade: 2, Preco: -500},
		{Nome: "Uva", Quantidade: 5, Preco: -750},
		{Nome: "Arroz", Quantidade: 1, Preco: 1000},
	}

	listaDeEmails := []string{
		"user1@gmail.com", "user2@gmail.com", "user3@gmail.com", "user4@gmail.com", "user5@gmail.com", "user6@gmail.com",
	}

	listaEmailEValores, isListaEmailEValoresValida := SepararValoresPorEmail(listaDeCompras, listaDeEmails)

	assertions.Empty(listaEmailEValores)
	assertions.Equal(false, isListaEmailEValoresValida)

}

// Separar Valores Por Email
// Lista De Emails Vazia
// Deve retornar uma lista de emails vazia
func TestSepararValoresPorEmail_ListaDeEmailsVazia(t *testing.T) {

	assertions := require.New(t)

	listaDeCompras := []model.Item{
		{Nome: "Detergente", Quantidade: 2, Preco: 500},
		{Nome: "Uva", Quantidade: 5, Preco: 750},
		{Nome: "Arroz", Quantidade: 1, Preco: 1000},
	}

	var listaDeEmails []string

	listaEmailEValores, isListaEmailEValoresValida := SepararValoresPorEmail(listaDeCompras, listaDeEmails)

	assertions.Empty(listaEmailEValores)
	assertions.Equal(false, isListaEmailEValoresValida)

}

// Separar Valores Por Email
// Lista De Compras Vazia
// Deve retornar uma lista de compras vazia
func TestSepararValoresPorEmail_ListaDeComprasVazia(t *testing.T) {

	assertions := require.New(t)

	var listaDeCompras []model.Item

	listaDeEmails := []string{
		"user1@gmail.com", "user2@gmail.com", "user3@gmail.com", "user4@gmail.com", "user5@gmail.com", "user6@gmail.com",
	}

	listaEmailEValores, isListaEmailEValoresValida := SepararValoresPorEmail(listaDeCompras, listaDeEmails)

	assertions.Empty(listaEmailEValores)
	assertions.Equal(false, isListaEmailEValoresValida)

}

// Separar Valores Por Email
// Lista De Compras E Emails Vazia
// Deve retornar uma lista de compras e emails vazia
func TestSepararValoresPorEmail_ListaDeComprasEEmailsVazia(t *testing.T) {

	assertions := require.New(t)

	var listaDeCompras []model.Item

	var listaDeEmails []string

	listaEmailEValores, isListaEmailEValoresValida := SepararValoresPorEmail(listaDeCompras, listaDeEmails)

	assertions.Empty(listaEmailEValores)
	assertions.Equal(false, isListaEmailEValoresValida)
}

// Separar Valores Por Email
// Valor Total Menor Que Quantidade De Emails
// Deve retornar uma lista de valores vazia
func TestSepararValoresPorEmail_ValorTotalMenorQueQuantidadeDeEmails(t *testing.T) {

	assertions := require.New(t)

	listaDeCompras := []model.Item{
		{Nome: "Detergente", Quantidade: 2, Preco: 2},
	}

	listaDeEmails := []string{
		"user1@gmail.com", "user2@gmail.com", "user3@gmail.com", "user4@gmail.com", "user5@gmail.com", "user6@gmail.com",
	}

	listaEmailEValores, isListaEmailEValoresValida := SepararValoresPorEmail(listaDeCompras, listaDeEmails)

	assertions.Empty(listaEmailEValores)
	assertions.Equal(false, isListaEmailEValoresValida)

}

// Separar Valores Por Email
// Lista De Compras Randomica
// Deve fazer a separação de valores por email usando uma lista de compras gerada randomicamente
func TestSepararValoresPorEmail_ListaDeComprasRandomica(t *testing.T) {

	assertions := require.New(t)

	listaDeCompras := util.GerarItens()

	listaDeEmails := []string{
		"user1@gmail.com", "user2@gmail.com", "user3@gmail.com", "user4@gmail.com", "user5@gmail.com",
	}

	listaEmailEValores, isListaEmailEValoresValida := SepararValoresPorEmail(listaDeCompras, listaDeEmails)

	// ! Pode gerar uma lista de compras vazia
	if isListaEmailEValoresValida.Status == false {
		assertions.NotEmpty(listaEmailEValores)
	} else {
		assertions.Empty(listaEmailEValores)
	}
}

// Separar Valores Por Email
// Lista De Emails Randomica
// Deve fazer a separação de valores por email usando uma lista de emails gerada randomicamente
func TestSepararValoresPorEmail_ListaDeEmailsRandomica(t *testing.T) {

	assertions := require.New(t)

	listaDeCompras := []model.Item{
		{Nome: "Detergente", Quantidade: 4, Preco: 500},
		{Nome: "Uva", Quantidade: 6, Preco: 750},
		{Nome: "Arroz", Quantidade: 2, Preco: 1000},
	}

	listaDeEmails := util.GerarEmails()

	listaEmailEValores, isListaEmailEValoresValida := SepararValoresPorEmail(listaDeCompras, listaDeEmails)

	// ! Pode gerar uma lista de emails vazia
	if isListaEmailEValoresValida.Status == false {
		assertions.NotEmpty(listaEmailEValores)
	} else {
		assertions.Empty(listaEmailEValores)
	}
}

// Separar Valores Por Email
// Lista De Compras E Emails Randomica
// Deve fazer a separação de valores por email usando uma lista de compras e emails gerada randomicamente
func TestSepararValoresPorEmail_ListaDeComprasEEmailsRandomica(t *testing.T) {

	assertions := require.New(t)

	listaDeCompras := util.GerarItens()

	listaDeEmails := util.GerarEmails()

	listaEmailEValores, isListaEmailEValoresValida := SepararValoresPorEmail(listaDeCompras, listaDeEmails)

	// ! Pode gerar uma lista de compras e/ou emails vazia
	if isListaEmailEValoresValida.Status == false {
		assertions.NotEmpty(listaEmailEValores)
	} else {
		assertions.Empty(listaEmailEValores)
	}
}
