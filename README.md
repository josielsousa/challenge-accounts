
# Challenge-Accounts

  

Desafio técnico em Go: API responsável por fornecer funcionalidades de criação de contas e transferências entre elas para um banco digital.

  

### Pré-requisitos

  

Possuir `docker` e `docker-compose` instalados.

  

-  [docker](https://docs.docker.com/get-docker/)

-  [docker-compose](https://docs.docker.com/compose/install/)

  

### Instalação

  

Após a instalação dos pré-requisitos, para disponibilizar a API, basta executar os próximos comandos:

```

git clone git@github.com:josielsousa/challenge-accounts.git

cd challenge-accounts

docker-compose up api

```

  

Executando com sucesso os comandos acima, a API estará disponível no endereço: http://localhost:3000/

  

## Tecnologias utilizadas

  

*  [GO](https://golang.org) - A linguagem usada

*  [docker](https://www.docker.com/)  e [docker-compose](https://www.docker.com/) - Ferramenta utilizada para criação da imagem base utilizada na construção da API. Ferramenta para definir e executar aplicativos `docker` de vários contêineres 

*  [MUX](github.com/gorilla/mux) - Framework utilizado para encapsular e facilitar a disponibilização dos métodos HTTP necessários para a API.

* [govalidator](github.com/thedevsaddam/govalidator)  - Framework utilizado para validação dos dados enviados na `request`.

*  [logrus](https://github.com/sirupsen/logrus) - Framework para formatação dos logs da API.

  *  [YAML](https://yaml.org/) - YAML é um formato de serialização de dados legíveis por humanos, utilizado para configurar os serviços da API para serem executados com o `docker-compose`.

* [uuid](github.com/google/uuid) - Utilizado na geração de universal ids.

*  [Gorm](github.com/jinzhu/gorm) - ORM para manipulação do acesso ao banco de dados

*  [sqlite3](https://www.sqlite.org/index.html) e [go-sqlite3](github.com/mattn/go-sqlite3) - Banco de dados em memória. Implementação do driver para acessar o banco de dados sqlite.

* [crypto](golang.org/x/crypto)  - Utilizado para gerar e validar o hash para autenticação do usuário.
* [decimal](github.com/shopspring/decimal) - Utilizado para realizar os cálculos nas transferências realizadas, mantendo assim a integridade do saldo nas contas. 
*  [jwt](https://jwt.io/) e [jwt-go/v4](github.com/dgrijalva/jwt-go/v4) - Utilizado na geração dos tokens para acesso as rotas privadas de transferências.

## Autor


*  **Josiel Sousa** - [josielsousa](https://github.com/josielsousa)