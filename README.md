## 💼 Desafio Técnico

## 📝 Sumário

- [Sobre](#about)
- [Estrutura](#pattern)
- [Dependências](#dep)
- [Como baixar e iniciar:](#howtodownload)
    - [Docker](#docker)
    - [Manual](#manual)
- [Como criar casos personalizados](#customcase)
- [Licença](#license)
- [Autor](#author)

## 💻 Sobre: <a name="about"></a>

- A aplicação divide o valor total resultante de uma lista de compras em uma lista de emails.

## 🏛 Estrutura: <a name="pattern"></a>

- A aplicação é composta das seguintes partes:
    - model: contém o arquivo <b>item.go</b>, que representa a entidade principal da aplicação.
    - service: contém o arquivo <b>item-service.go</b>, que contém as regras de serviço da entidade item.
    - util: contém os arquivos <b>gerador-de-emails.go</b> e <b>gerador-de-itens.go</b>, que auxiliam na criação de testes/casos para a aplicação.

## 🌴 Dependências: <a name="dep"></a>

- Testify: <a href="https://github.com/stretchr/testify">github.com/stretchr/testify</a>
- Randstr: <a href="https://github.com/thanhpk/randstr">github.com/thanhpk/randstr</a>

## 🔥 Como baixar e iniciar: <a name="howtodownload"></a>
## - 🐳 Docker: <a name="docker"></a>
⚠ Ter o <a href="https://www.docker.com/products/docker-desktop/">Docker</a> instalado.
- Clonar o repositório:
```bash
git clone https://github.com/Alberto-Pereira/Desafio-Tecnico.git
```
- Acessar o diretório clonado:
```bash
cd Desafio-Tecnico
```
- Fazer a build da aplicação:
```bash
docker build -t docker-desafio-tecnico .
```
- Iniciar a aplicação:
```bash
docker run docker-desafio-tecnico
```

## - 🔨 Manual: <a name="manual"></a>
- Clonar o repositório:
```bash
git clone https://github.com/Alberto-Pereira/Desafio-Tecnico.git
```
- Acessar o diretório clonado:
```bash
cd Desafio-Tecnico
```
- Baixar as dependências da aplicação:
```bash
go get .
```
- Iniciar a aplicação:
```bash
go run main.go
```

## 🧵 Como criar casos personalizados: <a name="customcase"></a>

- Copie o código e cole abaixo da <b>linha 177</b> do arquivo <b>main.go</b>:
```bash
listaDeCompras_NomeDoCaso := []model.Item{
    {Nome: "Goiaba", Quantidade: 6, Preco: 100},
    {Nome: "Pneu", Quantidade: 2, Preco: 30000},
    {Nome: "Lápis", Quantidade: 1, Preco: 200},
    {Nome: "Vassoura", Quantidade: 2, Preco: 500},
    {Nome: "Creme dental", Quantidade: 4, Preco: 250},
}

listaDeEmails_NomeDoCaso := []string{
    "user1@gmail.com", "user2@gmail.com", "user3@gmail.com", "user4@gmail.com",
    "user5@gmail.com", "user6@gmail.com", "user7@gmail.com", "user8@gmail.com",
}

listaEmailEValores_NomeDoCaso, isListaEmailEValoresValida_NomeDoCaso :=
    service.SepararValoresPorEmail(listaDeCompras_NomeDoCaso, listaDeEmails_NomeDoCaso)

if isListaEmailEValoresValida_NomeDoCaso {
    fmt.Println("\nLista de emails e valores:")
    fmt.Println(listaEmailEValores_NomeDoCaso)
}
```
- É possível alterar "_NomeDoCaso" para o caso que esteja criando.
- É possível alterar os valores da listaDeCompras e da listaDeEmails

## 📃 Licença: <a name="license"></a>

- <a href="http://www.apache.org/licenses/LICENSE-2.0.html">Apache 2.0</a>

## 👁‍🗨 Autor: <a name="author"></a>

- <a href="https://github.com/Alberto-Pereira">Alberto Pereira</a>
