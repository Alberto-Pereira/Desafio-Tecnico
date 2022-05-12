// Package util contém arquivos que auxiliam na criação de testes/casos para a aplicação.
package util

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/thanhpk/randstr"
)

// Gerar Emails
// Gera uma quantidade randômica de emails e retorna em uma lista
// É possível alterar algumas propriedades
func GerarEmails() []string {
	var listaDeEmails []string
	var emailsRepetidos int

	var quantidadeDeEmailsASerGerada int = 100 // Escolha até quantos emails serão gerados | 0 até n
	var quantidadeDeCaracteresDoEmail int = 10 // Escolha a quantidade de caracteres do email | Valor n será fixo

	seed := rand.NewSource(time.Now().UnixNano())                   // Gera uma seed
	random := rand.New(seed)                                        // Atribue a variável random um Rand usando a seed passada
	quantidadeDeEmails := random.Intn(quantidadeDeEmailsASerGerada) // Gera uma quantidade de emails randômica baseada na variável random

	// Cria uma lista com emails gerados randomicamente
	for i := 0; i < quantidadeDeEmails; i++ {

		listaDeEmails = append(listaDeEmails, (randstr.String(quantidadeDeCaracteresDoEmail) + "@gmail.com"))

	}

	// Checa por emails iguais
	// ! Na criação do map, chaves (emails) iguais são unificadas
	for i, isEmailUnico := range listaDeEmails {
		for j, emailASerComparado := range listaDeEmails {
			if i != j {
				if isEmailUnico == emailASerComparado {
					emailsRepetidos += 1
				}
			}
		}
	}

	if len(listaDeEmails) <= 0 {
		fmt.Println("Lista de emails gerada se encontra vazia!")
		return nil
	} else {
		fmt.Println("Quantidade de emails gerados:", quantidadeDeEmails)
	}

	if emailsRepetidos > 0 {
		fmt.Println("Emails repetidos:", emailsRepetidos)
		fmt.Println("Nova quantidade de emails gerados:", quantidadeDeEmails-(emailsRepetidos/2))
	}

	return listaDeEmails
}
