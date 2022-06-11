// Package principal da aplicação
package main

import (
	"desafiotecnico/model"
	"desafiotecnico/service"
	"desafiotecnico/util"
	"fmt"
)

// Main
// Executa alguns exemplos
// É possível adicionar casos personalizados ao final do código | Mais informações no repositório
func main() {

	// Exemplo - Módulo 0
	fmt.Println("\n- - Separar valores por email - Módulo 0 - -")

	listaDeCompras_Modulo0 := []model.Item{
		{Nome: "Goiaba", Quantidade: 6, Preco: 100},
		{Nome: "Pneu", Quantidade: 2, Preco: 30000},
		{Nome: "Lápis", Quantidade: 1, Preco: 200},
		{Nome: "Vassoura", Quantidade: 2, Preco: 500},
		{Nome: "Creme dental", Quantidade: 4, Preco: 250},
	}

	listaDeEmails_Modulo0 := []string{
		"user1@gmail.com", "user2@gmail.com", "user3@gmail.com", "user4@gmail.com",
		"user5@gmail.com", "user6@gmail.com", "user7@gmail.com", "user8@gmail.com",
	}

	listaEmailEValores_Modulo0, isListaEmailEValoresValida_Modulo0 :=
		service.SepararValoresPorEmail(listaDeCompras_Modulo0, listaDeEmails_Modulo0)

	if isListaEmailEValoresValida_Modulo0.Status == false {
		fmt.Println("\nLista de emails e valores:")
		fmt.Println(listaEmailEValores_Modulo0)
	}

	// Exemplo - Módulo diferente de 0
	fmt.Println("\n- - Separar valores por email - Módulo diferente de 0 - -")

	listaDeCompras_ModuloDiferenteDe0 := []model.Item{
		{Nome: "Laranja", Quantidade: 10, Preco: 150},
		{Nome: "Volante", Quantidade: 1, Preco: 25000},
		{Nome: "Caneta", Quantidade: 2, Preco: 350},
		{Nome: "Balde", Quantidade: 3, Preco: 1000},
		{Nome: "Escova", Quantidade: 5, Preco: 750},
	}

	listaDeEmails_ModuloDiferenteDe0 := []string{
		"user1@gmail.com", "user2@gmail.com", "user3@gmail.com",
	}

	listaEmailEValores_ModuloDiferenteDe0, isListaEmailEValoresValida_ModuloDiferenteDe0 :=
		service.SepararValoresPorEmail(listaDeCompras_ModuloDiferenteDe0, listaDeEmails_ModuloDiferenteDe0)

	if isListaEmailEValoresValida_ModuloDiferenteDe0.Status == false {
		fmt.Println("\nLista de emails e valores:")
		fmt.Println(listaEmailEValores_ModuloDiferenteDe0)
	}

	// Exemplo - Lista de compras vazia
	fmt.Println("\n- - Separar valores por email - Lista de compras vazia - -")

	var listaDeCompras_ListaDeComprasVazia []model.Item

	listaDeEmails_ListaDeComprasVazia := []string{
		"user1@gmail.com", "user2@gmail.com", "user3@gmail.com", "user4@gmail.com",
		"user5@gmail.com", "user6@gmail.com", "user7@gmail.com", "user8@gmail.com",
	}

	listaEmailEValores_ListaDeComprasVazia, isListaEmailEValoresValida_ListaDeComprasVazia :=
		service.SepararValoresPorEmail(listaDeCompras_ListaDeComprasVazia, listaDeEmails_ListaDeComprasVazia)

	if isListaEmailEValoresValida_ListaDeComprasVazia.Status == false {
		fmt.Println("\nLista de emails e valores:")
		fmt.Println(listaEmailEValores_ListaDeComprasVazia)
	}

	// Exemplo - Lista de emails vazia
	fmt.Println("\n- - Separar valores por email - Lista de emails vazia - -")

	listaDeCompras_ListaDeEmailsVazia := []model.Item{
		{Nome: "Goiaba", Quantidade: 6, Preco: 100},
		{Nome: "Pneu", Quantidade: 2, Preco: 30000},
		{Nome: "Lápis", Quantidade: 1, Preco: 200},
		{Nome: "Vassoura", Quantidade: 2, Preco: 500},
		{Nome: "Creme dental", Quantidade: 4, Preco: 250},
	}

	var listaDeEmails_ListaDeEmailsVazia []string

	listaEmailEValores_ListaDeEmailsVazia, isListaEmailEValoresValida_ListaDeEmailsVazia :=
		service.SepararValoresPorEmail(listaDeCompras_ListaDeEmailsVazia, listaDeEmails_ListaDeEmailsVazia)

	if isListaEmailEValoresValida_ListaDeEmailsVazia.Status == false {
		fmt.Println("\nLista de emails e valores:")
		fmt.Println(listaEmailEValores_ListaDeEmailsVazia)
	}

	// Exemplo - Lista de compras e emails vazia
	fmt.Println("\n- - Separar valores por email - Lista de compras e emails vazia - -")

	var listaDeCompras_ListaDeComprasEEmailsVazia []model.Item

	var listaDeEmails_ListaDeComprasEEmailsVazia []string

	listaEmailEValores_ListaDeComprasEEmailsVazia, isListaEmailEValoresValida_ListaDeComprasEEmailsVazia :=
		service.SepararValoresPorEmail(listaDeCompras_ListaDeComprasEEmailsVazia, listaDeEmails_ListaDeComprasEEmailsVazia)

	if isListaEmailEValoresValida_ListaDeComprasEEmailsVazia.Status == false {
		fmt.Println("\nLista de emails e valores:")
		fmt.Println(listaEmailEValores_ListaDeComprasEEmailsVazia)
	}

	// Exemplo - Lista de compras gerada randomicamente
	// ! É possível que uma lista de compras vazia seja gerada
	fmt.Println("\n- - Separar valores por email - Lista de compras gerada randomicamente - -")

	listaDeCompras_ListaDeComprasRandom := util.GerarItens()

	listaDeEmails_ListaDeComprasRandom := []string{
		"user1@gmail.com", "user2@gmail.com", "user3@gmail.com", "user4@gmail.com",
	}

	listaEmailEValores_ListaDeComprasRandom, isListaEmailEValoresValida_ListaDeComprasRandom :=
		service.SepararValoresPorEmail(listaDeCompras_ListaDeComprasRandom, listaDeEmails_ListaDeComprasRandom)

	if isListaEmailEValoresValida_ListaDeComprasRandom.Status == false {
		fmt.Println("\nLista de emails e valores:")
		fmt.Println(listaEmailEValores_ListaDeComprasRandom)
	}

	// Exemplo - Lista de emails gerada randomicamente
	// ! É possível que uma lista de emails vazia seja gerada
	fmt.Println("\n- - Separar valores por email - Lista de emails gerada randomicamente - -")

	listaDeCompras_ListaDeEmailsRandom := []model.Item{
		{Nome: "Goiaba", Quantidade: 6, Preco: 100},
		{Nome: "Pneu", Quantidade: 2, Preco: 30000},
		{Nome: "Lápis", Quantidade: 1, Preco: 200},
		{Nome: "Vassoura", Quantidade: 2, Preco: 500},
		{Nome: "Creme dental", Quantidade: 4, Preco: 250},
	}

	listaDeEmails_ListaDeEmailsRandom := util.GerarEmails()

	listaEmailEValores_ListaDeEmailsRandom, isListaEmailEValoresValida_ListaDeEmailsRandom :=
		service.SepararValoresPorEmail(listaDeCompras_ListaDeEmailsRandom, listaDeEmails_ListaDeEmailsRandom)

	if isListaEmailEValoresValida_ListaDeEmailsRandom.Status == false {
		fmt.Println("\nLista de emails e valores:")
		fmt.Println(listaEmailEValores_ListaDeEmailsRandom)
	}

	// Exemplo - Lista de compras e emails gerada randomicamente
	// ! É possível que uma lista de compras e/ou emails vazia seja gerada
	fmt.Println("\n- - Separar valores por email - Lista de compras e emails gerada randomicamente - -")

	listaDeCompras_ListaDeComprasEEmailsRandom := util.GerarItens()

	listaDeEmails_ListaDeComprasEEmailsRandom := util.GerarEmails()

	listaEmailEValores_ListaDeComprasEEmailsRandom, isListaEmailEValoresValida_ListaDeComprasEEmailsRandom :=
		service.SepararValoresPorEmail(listaDeCompras_ListaDeComprasEEmailsRandom, listaDeEmails_ListaDeComprasEEmailsRandom)

	if isListaEmailEValoresValida_ListaDeComprasEEmailsRandom.Status == false {
		fmt.Println("\nLista de emails e valores:")
		fmt.Println(listaEmailEValores_ListaDeComprasEEmailsRandom)
	}

	fmt.Println(" - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")

	// Casos Personalizados
	// Abaixo você pode colar o código informado no repositório e alterar suas propriedades
	fmt.Println("\n- - Caso Personalizado - Separar valores por email - Nome do caso - -")
}
