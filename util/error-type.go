// Package util contém arquivos que auxiliam no desenvolvimento da aplicação.
package util

// ErrorType é usado para lidar com possíveis erros e/ou valores vazios
type ErrorType struct {
	Mensagem string // Mensagem do erro
	Status   bool   // Status do erro
}
