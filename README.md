# Challenge-Accounts

Desafio técnico em Go: API responsável por fornecer funcionalidades de criação de contas e transferências entre elas para um banco digital.

### Pré-requisitos

Possuir `docker` e `docker-compose` instalados.

-  [docker](https://docs.docker.com/get-docker/)

-  [docker-compose](https://docs.docker.com/compose/install/)

### Instalação

Após a instalação dos pré-requisitos, para disponibilizar a API, basta executar os próximos comandos:

```bash
git clone git@github.com:josielsousa/challenge-accounts.git
cd challenge-accounts
docker-compose up api
```

### Testes unitários

A API foi construída possuindo testes unitários, para executar os testes, basta executar os seguintes comandos:

```bash
git clone git@github.com:josielsousa/challenge-accounts.git
cd challenge-accounts
docker-compose up api-test
```

  

  

Após a execução de todos testes unitários, será criado um arquivo chamado `coverage.html`, contendo o resumo da cobertura dos testes realizados.

  
  

### Rotas / Endpoints

  

*  `/accounts` - `POST` - Rota utilizada para criação de uma nova `account`

	Payload de entrada:
	```json
	{
		"cpf": "47298817027",
		"name": "Joãozinho",
		"secret": "secret",
		"balance": 80.99
	}
	```
	| Atributo| Obrigatório | Descrição
	|--|--|--|
	| name | SIM | Nome do usuário que a conta que será criada.
	| cpf | SIM | Número de CPF do usuário. Deve ser utilizado um número válido
	| secret | SIM | Senha para ser utilizada na autenticação do usuário
	| balance | NÃO | Saldo inicial da conta que será criada.


Payload Retornos: 

*  Status Code `200` - Sucesso, conta criada 
	```json
	{
		"id": "ac23d91c-08b0-45aa-90d9-9534207c318e",
		"cpf": "47298817027",
		"name": "Joãozinho",
		"secret": "$2a$10$ID9bqUy9DvXqKGrGsRuJBuiZ.WvTcbs9X.UkEEHcYuUdu5IuULtNm",
		"balance": 80.99,
		"created_at": "2020-07-26T21:25:41.123199754Z"
	}
	```

	
	| Atributo| Descrição
	|--|--|
	| id | UUID gerado para a `account`.
	| name | Nome do usuário.
	| cpf | Número de CPF do usuário. 
	| secret | Hash gerado para o `secret` informado.
	| balance | Saldo inicial da conta que será criada.
	| created_at | Data de criação da conta 


* Status Code `422` -  Erro - Os dados de entrada são válidos porém existe uma `account` para o CPF informado. 
	```json
	{
	  "error": "Já existe uma conta criada com o CPF informado."
	}
	```

* Status Code `500` -  Erro inesperado durante o processamento da requisição
	```json
	{
	  "error": "Erro Inesperado"
	}
	```
---

*  `/accounts` - `GET` - Rota utilizada para listagem de todas as `accounts`


* Status Code `204` -  Requisição executada com sucesso, porém não possui dados de retorno, lista de `accounts` vazia.
		
	Payload de retorno: Não possui

* Status Code `200` -  Quando existir accounts para serem retornadas

	Payload de retorno:
	```json
	{
	  "data": [
	    {
	      "id": "011bd273-6f88-4488-8d69-abee5c41340f",
	      "cpf": "47298817027",
	      "name": "Joãozinho",
	      "secret": "$2a$10$FNcpXUGL1nhnF51t2Xojt.B.dndOKfB9zZOyy3n5hdF9cg3gnFZMq",
	      "balance": 80.99,
	      "created_at": "2020-07-26T23:04:30.904920793Z"
	    },
	    {
	      "id": "689b2629-f02a-412e-873f-cfaab344b413",
	      "cpf": "14394183901",
	      "name": "Emanuelly Cláudia Jennifer",
	      "secret": "$2a$10$bKLye7aAf/f.D/tgHH4vDuo6KDK17mKFM.Thmt/aSLFzJTP4Bndny",
	      "balance": 0,
	      "created_at": "2020-07-26T23:05:01.010789798Z"
	    }
	  ],
	  "success": true
	}
	```
	
	| Atributo| Descrição
	|--|--|
	| data | Lista de `accounts`
	| success | Boolean - Indica o status de sucesso para a requisição

* Status Code `500` -  Erro inesperado durante o processamento da requisição
	```json
	{
	  "error": "Erro Inesperado"
	}
	```
---

*  `/accounts/{id}/balance` - `GET` - Rota utilizada para recuperar o saldo de uma  `account`

	Payload de retorno:
	```json
	{
		"data": {
		    "balance": 80.99
		},
		"success": true
	}
	```
	
	| Atributo| Descrição
	|--|--|
	| data | Retorna o saldo da `account` no atributo `balance`
	| success | Boolean - Indica o status de sucesso para a requisição

* Status Code `500` -  Erro inesperado durante o processamento da requisição
	```json
	{
	  "error": "Erro Inesperado"
	}
	```

* Status Code `404` -  Quando não encontrar a account conforme o `id` informado
	```json
	{
	  "error": "Conta não encontrada"
	}
	```

---


### Tecnologias utilizadas

*  [GO](https://golang.org) - A linguagem usada

*  [docker](https://www.docker.com/) e [docker-compose](https://www.docker.com/) - Ferramenta utilizada para criação da imagem base utilizada na construção da API. Ferramenta para definir e executar aplicativos `docker` de vários contêineres

*  [MUX](github.com/gorilla/mux) - Framework utilizado para encapsular e facilitar a disponibilização dos métodos HTTP necessários para a API.

*  [govalidator](github.com/thedevsaddam/govalidator) - Framework utilizado para validação dos dados enviados na `request`.

*  [logrus](https://github.com/sirupsen/logrus) - Framework para formatação dos logs da API.

*  [YAML](https://yaml.org/) - YAML é um formato de serialização de dados legíveis por humanos, utilizado para configurar os serviços da API para serem executados com o `docker-compose`.

*  [uuid](github.com/google/uuid) - Utilizado na geração de universal ids.

*  [Gorm](github.com/jinzhu/gorm) - ORM para manipulação do acesso ao banco de dados

*  [sqlite3](https://www.sqlite.org/index.html) e [go-sqlite3](github.com/mattn/go-sqlite3) - Banco de dados em memória. Implementação do driver para acessar o banco de dados sqlite.

*  [crypto](golang.org/x/crypto) - Utilizado para gerar e validar o hash para autenticação do usuário.

*  [decimal](github.com/shopspring/decimal) - Utilizado para realizar os cálculos nas transferências realizadas, mantendo assim a integridade do saldo nas contas.

*  [jwt](https://jwt.io/) e [jwt-go/v4](github.com/dgrijalva/jwt-go/v4) - Utilizado na geração dos tokens para acesso as rotas privadas de transferências.

### Autor

*  **Josiel Sousa** - [josielsousa](https://github.com/josielsousa)