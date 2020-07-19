# Challenge-Accounts

Desafio técnico em Go:  API  responsável por fornecer funcionalidades de criação de contas e transferências entre elas para um banco digital.

### Pré-requisitos

Possuir `docker` e `docker-compose` instalados.

 - [docker](https://docs.docker.com/get-docker/)
 - [docker-compose](https://docs.docker.com/compose/install/)

### Instalação

Após a instalação dos pré-requisitos, para disponibilizar a API, basta executar os próximos comandos: 
```
git clone git@github.com:josielsousa/challenge-accounts.git
cd challenge-accounts
docker-compose up api
```

Executando com sucesso os comandos acima, a API estará disponível no endereço: http://localhost:3000/

## Tecnologias utilizadas

* [GO](https://golang.org) - A linguagem usada

* [MUX](github.com/gorilla/mux) - Framework utilizado para encapsular e facilitar a disponibilização dos métodos HTTP necessários para a API.

* [logrus](https://github.com/sirupsen/logrus) - Framework para formatação dos logs da API.

* [docker](https://www.docker.com/) - Ferramenta utilizada para criação da imagem base utilizada na construção da API.

* [docker-compose](https://www.docker.com/) - Ferramenta para definir e executar aplicativos `docker` de vários contêineres
 
* [YAML](https://yaml.org/) - YAML é um formato de serialização de dados legíveis por humanos, utilizado para configurar os serviços da API para serem executados com o `docker-compose`.

* [Gorm](github.com/jinzhu/gorm) - ORM para manipulação do acesso ao banco de dados

* [sqlite3](https://www.sqlite.org/index.html) - Banco de dados em memória.

* [go-sqlite3](github.com/mattn/go-sqlite3) - Implementação do driver para acessar o banco de dados sqlite.

## Autor

* **Josiel Sousa** - [josielsousa](https://github.com/josielsousa)