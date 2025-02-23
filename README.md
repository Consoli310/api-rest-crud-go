# API RESTful em Go

Esta é uma API RESTful desenvolvida com **Go** e **Gorilla Mux** para gerenciamento de produtos. A API permite operações CRUD (Create, Read, Update, Delete) utilizando **MySQL** como banco de dados.

## 🚀 Tecnologias Utilizadas

- **Linguagem:** Go
- **Framework:** Gorilla Mux
- **Banco de Dados:** MySQL
- **Gerenciamento de Dependências:** Go Modules

## 📌 Funcionalidades

- Cadastro de produtos com os seguintes campos:
  - Nome
  - Descrição
  - Preço
  - Quantidade
- API RESTful para operações CRUD
- Integração com banco de dados MySQL

## ⚙️ Como Executar o Projeto

1. Clone o repositório:
   ```sh
   git clone https://github.com/Consoli310/api-rest-crud-go.git
   ```
2. Acesse o diretório do projeto:
   ```sh
   cd api-rest-crud-go
   ```
3. Instale as dependências:
   ```sh
   go mod download
   ```
4. Configure o banco de dados MySQL:
   - Crie um banco de dados chamado `productsdb`
   - Atualize as credenciais no arquivo de configuração

5. Execute a aplicação:
   ```sh
   go run main.go
   ```
6. A API estará disponível em:
   ```
   http://localhost:8080
   ```

## 📌 Rotas da API

- **POST** `/products` - Criar um novo produto
- **GET** `/products` - Listar todos os produtos
- **GET** `/products/{id}` - Buscar um produto pelo ID
- **PUT** `/products/{id}` - Atualizar um produto pelo ID
- **DELETE** `/products/{id}` - Remover um produto pelo ID

## 🗄️ Banco de Dados

A API utiliza **MySQL** para armazenar os produtos cadastrados. Certifique-se de configurar corretamente as credenciais de acesso no código antes de executar a aplicação.

