# go-qr-code-generator

API escrita em Go que gera imagens QR Code baseado nas entradas do usuário.

### 📋 Pré-requisitos

Antes de começar, você vai precisar ter instalado em sua máquina a linguagem Go. As instruções estão disponíveis neste [link](https://go.dev/doc/install). Você pode verificar a instalação digitando o comando:

```
which go
```

Em seu terminal. A resposta deverá ser o path onde o compilador se encontra.

Também será necessário ter instalado o sistema de controle de versões Git. As instruções de instalação estão neste [link](https://git-scm.com/book/pt-br/v2/Come%C3%A7ando-Instalando-o-Git)

## 🚀 Começando

- Clone o repositório **go-qr-code-generator** em seu computador:

```
git clone git@github.com:ddvalim/go-qr-code-generator.git
```

- Instale as dependências do projeto:

```
go get ./...
```

## 🔧 Execução

Para executar o servidor:

```
go run main.go
```

A API estará executando na porta 8080.

## 📍 Rotas

- Gerar um QR Code:

```
curl --request GET \
  --url http://localhost:8080/qrcode \
  --header 'Content-Type: application/json' \
  --data '{
	"content": "hello, world!",
	"image_size": 256,
	"version_number": 1,
	"level": 1
}'
```

## 🛠️ Construído com

* [go-qrcode](https://pkg.go.dev/github.com/skip2/go-qrcode@v0.0.0-20200617195104-da1b6568686e) - Biblioteca utilizada para gerar imagens QR Code
* [Mux](https://pkg.go.dev/github.com/gorilla/mux) - Request router

## ✒️ Autora

Este projeto foi desenvolvido por Diovana Rodrigues Valim, bacharela em Sistemas de Informação pela Universidade Federal de Santa Catarina & Engenheira de Software no Mercado Livre.

---
