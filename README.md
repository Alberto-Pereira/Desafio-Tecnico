## 💼 Desafio Técnico

## 📝 Sumário

- [Sobre](#about)
- [Estrutura](#pattern)
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

## 🔥 Como baixar e iniciar: <a name="howtodownload"></a>
## - 🐳 Docker: <a name="docker"></a>
⚠ Ter o <a href="https://www.docker.com/products/docker-desktop/">Docker</a> instalado.
- Clonar o repositório:
```bash
git clone https://github.com/Alberto-Pereira/Desafio-Tecnico.git
```
⚠ Acesse o diretório clonado:
```bash
cd MinhasCriptos-API
```
- Fazer a build da aplicação:
```bash
docker-compose build
```
- Iniciar a aplicação:
```bash
docker-compose up
```

## - 🔨 Manual: <a name="manual"></a>

## 🧵 Como criar casos personalizados: <a name="customcase"></a>

## 📃 Licença: <a name="license"></a>

- <a href="http://www.apache.org/licenses/LICENSE-2.0.html">Apache 2.0</a>

## 👁‍🗨 Autor: <a name="author"></a>

- <a href="https://github.com/Alberto-Pereira">Alberto Pereira</a>
