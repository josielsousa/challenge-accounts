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
make setup-dev
make run
```

---

### Testes unitários

A API foi construída possuindo testes unitários, para executar os testes, basta executar os seguintes comandos:

```bash
git clone https://github.com/josielsousa/challenge-accounts.git
cd challenge-accounts
docker-compose up api-test
```
