# syntax=docker/dockerfile:1

# Versão escolhida do golang
FROM golang:1.18

# Diretório da imagem
WORKDIR /app

# Copiar os arquivos que contém as dependências
COPY go.mod ./
COPY go.sum ./

# Baixar dependências
RUN go mod download

# Copiar todos os arquivos do app
COPY . .

# Fazer a build do app
RUN go build -o /docker-desafio-tecnico

# Executar
CMD ["/docker-desafio-tecnico"]