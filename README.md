# Challenge-Accounts

Desafio técnico em Go: API responsável por fornecer funcionalidades de criação de contas e transferências entre elas para um banco digital.

### Pré-requisitos

Possuir `docker` e `docker-compose` instalados.

-  [docker](https://docs.docker.com/get-docker/)

-  [docker-compose](https://docs.docker.com/compose/install/)

---

### Instalação

Após a instalação dos pré-requisitos, para disponibilizar a API, basta executar os próximos comandos:

```bash
git clone https://github.com/josielsousa/challenge-accounts.git
cd challenge-accounts
docker-compose up api
```

---

### Testes unitários

A API foi construída possuindo testes unitários, para executar os testes, basta executar os seguintes comandos:

```bash
git clone https://github.com/josielsousa/challenge-accounts.git
cd challenge-accounts
docker-compose up api-test
```


Após a execução de todos testes unitários, será criado um arquivo chamado `coverage.html`, contendo o resumo da cobertura dos testes realizados.

  ---
  
### Rotas / Endpoints - `Accounts`


*  `/accounts` - `POST` - Rota utilizada para criação de uma nova `account`


	Exemplo de requisição : 
	```bash 
	curl --request POST -v \
	  --url http://localhost:3000/accounts \
	  --header 'content-type: application/json' \
	  --data '{
		"cpf": "04075532151",
		"name": "Juliana Alice Luana Nunes",
		"secret": "secret",
		"balance": 99.50
	}'
	```


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


	Exemplo de requisição : 
	```bash 
	curl --request GET -v --url http://localhost:3000/accounts
	```

* Status Code `204` -  Requisição executada com sucesso, porém não possui dados de retorno, lista de `accounts` vazia.
		
	Payload de retorno: Não se aplica.

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
	    },
	    {
	      "id": "41a721e6-16d3-45ae-8604-b6239fea31ae",
	      "cpf": "04075532151",
	      "name": "Juliana Alice Luana Nunes",
	      "secret": "$2a$10$e3hj/3mVf4.K8cb.s50aKOrN2SXF2ZWYLJeigev1hJCjOXxYI68te",
	      "balance": 99.5,
	      "created_at": "2020-07-26T23:47:48.877912737Z"
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


	Exemplo de requisição : 
	```bash 
	curl --request GET -v \
	  --url http://localhost:3000/accounts/{id}/balance
	```

	Payload de retorno:

*  Status Code `200` - Sucesso, conta encontrada, saldo retornado: 

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
  
### Rotas / Endpoints - `Login`


*  `/login` - `POST` - Rota utilizada para autenticação de usuário que possua `account`.

	Exemplo de requisição : 
	```bash 
	curl --request POST -v \
	  --url http://localhost:3000/login \
	  --header 'content-type: application/json' \
	  --data '{
		"cpf": "04075532151",
		"secret": "secret"
	}'
	```

	Payload de entrada:
	```json
	{
		"cpf": "14394183901",
		"secret": "secret2"
	}
	```
	| Atributo| Obrigatório | Descrição
	|--|--|--|
	| cpf | SIM | Número de CPF do usuário. Deve ser utilizado um número válido
	| secret | SIM | Senha para ser utilizada na autenticação do usuário, deve ser a mesma informada na criação da `account`


	Payload Retornos: 

*  Status Code `200` - Quando a autenticação for bem sucedida.
	```json
	{
	  "data": {
	    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjE0Mzk0MTgzOTAxIiwiYWNjb3VudF9pZCI6IjY4OWIyNjI5LWYwMmEtNDEyZS04NzNmLWNmYWFiMzQ0YjQxMyIsImV4cGlyZXNfYXQiOjE1OTU4MDU4NzZ9.TtJjppQasSjNlpR6Y_ljeA9wWzFzCSKHc8RuWaKT3lw"
	  },
	  "success": true
	}
	```

	| Atributo| Descrição
	|--|--|
	| data | Retorna o `token` _JWT_ para acesso as rotas privadas.
	| success | Boolean - Indica o status de sucesso para a requisição

* Status Code `401` - Quando o `secret` fornecido for diferente do `secret` armazenado.
	```json
	{
	  "error": "Não autenticado"
	}
	```

* Status Code `404`  - Quando não encontrar a account.
	```json
	{
	  "error": "Conta não encontrada"
	}
	```

* Status Code `500` -  Erro inesperado durante o processamento da requisição
	```json
	{
	  "error": "Erro Inesperado"
	}
	```

---
  
### Rotas / Endpoints - `Transfers`

O Header `Access-Token` é obrigatório para utilização dos `endpoints` de transferência, deve ser um `token` válido, obtido através do endpoint `/login`.  Esses são os possíveis retornos durante a validação do `token` informado: 


* Status Code `400` -  Quando o `token` estiver vazio / nulo
	```json
	{
	  "error": "Token vazio."
	}
	```

* Status Code `401` -  Quando o `token` fornecido for inválido.
	```json
	{
	  "error": "A chave de assinatura do token é inválida."
	}
	```

* Status Code `401` -  Quando o `token` estiver expirado.
	```json
	{
	  "error": "Token expirado."
	}
	```

* Status Code `500` -  Erro inesperado durante o processamento da requisição
	```json
	{
	  "error": "Erro Inesperado"
	}
	```
---

*  `/transfers` - `POST` - Realiza a transferência entre as `accounts`

	Exemplo de requisição : 
	```bash 
	curl --request POST -v \
	  --url http://localhost:3000/transfers \
	  --header 'access-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjA0MDc1NTMyMTUxIiwiYWNjb3VudF9pZCI6IjQxYTcyMWU2LTE2ZDMtNDVhZS04NjA0LWI2MjM5ZmVhMzFhZSIsImV4cGlyZXNfYXQiOjE1OTU4MDk4MTh9.NAA7z3TQW5qD_P6Gl92vcfMuFba5J4k-LI-iLiWb5x4' \
	  --header 'content-type: application/json' \
	  --data '{
		"account_destination_id": "689b2629-f02a-412e-873f-cfaab344b413",
		"amount": 0.33
	}'
	```


	Payload de entrada:

	```json
	{
		"account_destination_id": "689b2629-f02a-412e-873f-cfaab344b413",
		"amount": 0.01
	}
	```
	| Atributo| Obrigatório | Descrição
	|--|--|--|
	| account_destination_id | SIM | Identificador universal UUID da conta de destino do `amount`
	| amount | SIM | Valor que será transferido da `account_origin_id` para a `account_destination_id` 


	Payload Retornos: 

*  Status Code `200` - Quando a transferência for bem sucedida.
	```json
	{
	  "data": {
	    "id": "ae4c2025-ac57-48ab-a5ed-b9266fe52314",
	    "account_origin_id": "41a721e6-16d3-45ae-8604-b6239fea31ae",
	    "account_destination_id": "689b2629-f02a-412e-873f-cfaab344b413",
	    "amount": 0.01,
	    "created_at": "2020-07-26T23:48:58.858125762Z"
	  },
	  "success": true
	}
	```

	| Atributo| Descrição
	|--|--|
	| data | Retorna os dados da transferência realizada
	| id | Identificador universal UUID da transferência realizada.
	| account_origin_id | Identificador universal UUID da conta de origem do `amount`
	| account_destination_id | Identificador universal UUID da conta de destino do `amount`
	| amount | Valor que foi transferido da `account_origin_id` para a `account_destination_id` 
	| created_at | Data de realização da transferência 
	| success | Boolean - Indica o status de sucesso para a requisição


* Status Code `404` -  Quando a `account` origem não for encontrada
	A conta de origem será recuperada do token informado.
	```json
	{
	  "error": "Conta de origem não encontrada"
	}
	```

* Status Code `404` -  Quando a `account` destino não for encontrada
	```json
	{
	  "error": "Conta de destino não encontrada"
	}
	```

* Status Code `422` -  Quando não houver saldo disponível suficiente na `account` de origem.
	```json
	{
	 "error": "Conta de origem sem saldo disponível"
	}
	```
---

*  `/transfers` - `GET` - Recupera todas as transferências realizadas para o usuário autenticado, a `account` será recuperar a através do `token` informado no header `Access-Token`

	Exemplo de requisição : 
	```bash 
	curl --request GET -v \
	  --url http://localhost:3000/transfers \
	  --header 'access-token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IjA0MDc1NTMyMTUxIiwiYWNjb3VudF9pZCI6IjQxYTcyMWU2LTE2ZDMtNDVhZS04NjA0LWI2MjM5ZmVhMzFhZSIsImV4cGlyZXNfYXQiOjE1OTU4MDk4MTh9.NAA7z3TQW5qD_P6Gl92vcfMuFba5J4k-LI-iLiWb5x4' \
	  --header 'content-type: application/json'
	```

	Payload de entrada: Não se aplica, a `account` será obtida através do `token` informado.


	Payload Retornos: 

*  Status Code `200` - Quando existir transferências a serem retornadas.
	```json
	{
	  "data": [
	    {
	      "id": "ae4c2025-ac57-48ab-a5ed-b9266fe52314",
	      "account_origin_id": "41a721e6-16d3-45ae-8604-b6239fea31ae",
	      "account_destination_id": "689b2629-f02a-412e-873f-cfaab344b413",
	      "amount": 0.01,
	      "created_at": "2020-07-26T23:48:58.858125762Z"
	    },
	    {
	      "id": "a854fd37-3770-4a64-b8d0-d47321ccd870",
	      "account_origin_id": "41a721e6-16d3-45ae-8604-b6239fea31ae",
	      "account_destination_id": "689b2629-f02a-412e-873f-cfaab344b413",
	      "amount": 0.01,
	      "created_at": "2020-07-26T23:57:31.34448352Z"
	    },
	    {
	      "id": "23000c5b-a8a5-4380-934a-a24afa79ecc2",
	      "account_origin_id": "41a721e6-16d3-45ae-8604-b6239fea31ae",
	      "account_destination_id": "689b2629-f02a-412e-873f-cfaab344b413",
	      "amount": 0.33,
	      "created_at": "2020-07-27T00:25:34.391882336Z"
	    }
	  ],
	  "success": true
	}
	```

	| Atributo| Descrição
	|--|--|
	| data | Retorna uma lista com todas as transferências realizadas
	| success | Boolean - Indica o status de sucesso para a requisição


* Status Code `204` -  Requisição executada com sucesso, porém não possui dados de retorno, lista de `transfers` vazia.
		
	Payload de retorno: Não se aplica.
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

* [curl](https://curl.haxx.se/) - Utilizado nos exemplos de requisição dessa documentação.

* [4devs - Gerador de Pessoas](https://www.4devs.com.br/gerador_de_pessoas) - Mock utilizado para gerar dados válidos durante testes e documentação da API.

---
### Autor

*  **Josiel Sousa** - [josielsousa](https://github.com/josielsousa)


