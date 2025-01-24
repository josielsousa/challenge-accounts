// Package swag Code generated by swaggo/swag. DO NOT EDIT
package swag

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/challenge-accounts/accounts": {
            "get": {
                "description": "Endpoint utilizado listar todas as contas.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Lista as contas.",
                "responses": {
                    "200": {
                        "description": "Lista todas as contas",
                        "schema": {
                            "$ref": "#/definitions/ListAccountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Endpoint utilizado para criação de uma nova conta.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Criar nova conta.",
                "parameters": [
                    {
                        "description": "Dados da conta",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.CreateAccountRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Conta criada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/CreateAccountResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/challenge-accounts/accounts/{account_id}/balance": {
            "get": {
                "description": "Endpoint utilizado consultar o saldo da conta.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Retorna o saldo da conta.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador da conta de origem",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Saldo da conta.",
                        "schema": {
                            "$ref": "#/definitions/GetAccountBalanceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Conta não encontrada.",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/challenge-accounts/accounts/{account_id}/transfers": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Endpoint utilizado para listar as transferências entre contas.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Lista as transferência da conta.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Identificador da conta de origem",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Transferência realizada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/ListTransfersResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Conta não encontrada.",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Endpoint utilizado para realizar uma transferência entre contas.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Realizar transferência entre contas.",
                "parameters": [
                    {
                        "description": "Dados da transação",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/DoTransferRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Identificador da conta de origem",
                        "name": "account_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Transferência realizada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/TransferResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Conta não encontrada.",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/challenge-accounts/auth/signin": {
            "post": {
                "description": "Endpoint utilizado para autorizar o acesso a conta.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Autoriza o acesso a conta.",
                "parameters": [
                    {
                        "description": "Dados de identificação",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/SinginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Conta criada com sucesso.",
                        "schema": {
                            "$ref": "#/definitions/SinginReponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateAccountResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "description": "Balance - saldo atual da conta.",
                    "type": "integer"
                },
                "cpf": {
                    "description": "CPF - número do CPF do titular da conta.",
                    "type": "string"
                },
                "created_at": {
                    "description": "CreatedAt - data de criação da conta.",
                    "type": "string"
                },
                "id": {
                    "description": "ID - identificador único da conta.",
                    "type": "string"
                },
                "name": {
                    "description": "Name - nome do titular da conta.",
                    "type": "string"
                }
            }
        },
        "DoTransferRequest": {
            "type": "object",
            "required": [
                "account_destination_id",
                "amount"
            ],
            "properties": {
                "account_destination_id": {
                    "description": "AccountDestinationID - identificador da conta de destino.",
                    "type": "string"
                },
                "amount": {
                    "description": "Amount - valor a ser transferido.",
                    "type": "integer"
                }
            }
        },
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "GetAccountBalanceResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "description": "Balance - saldo atual da conta.",
                    "type": "integer"
                }
            }
        },
        "ListAccountResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Data - lista de contas.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/CreateAccountResponse"
                    }
                },
                "success": {
                    "description": "Succes - indica se a operação foi bem sucedida.",
                    "type": "boolean"
                }
            }
        },
        "ListTransfersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Data - lista de transferências.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/TransferResponse"
                    }
                },
                "success": {
                    "description": "Succes - indica se a operação foi bem sucedida.",
                    "type": "boolean"
                }
            }
        },
        "SinginReponse": {
            "type": "object",
            "properties": {
                "token": {
                    "description": "Token - token de autenticação usado para acessar endpoints privados.",
                    "type": "string"
                }
            }
        },
        "SinginRequest": {
            "description": "Dados de identificação para o signin",
            "type": "object",
            "required": [
                "cpf",
                "password"
            ],
            "properties": {
                "cpf": {
                    "description": "Cpf - número do CPF do titular da conta, usado como username para o signin.",
                    "type": "string"
                },
                "password": {
                    "description": "Password - senha de acesso a conta.",
                    "type": "string"
                }
            }
        },
        "TransferResponse": {
            "type": "object",
            "properties": {
                "account_destination_id": {
                    "description": "AccountDestinationID - identificador da conta de destino.",
                    "type": "string"
                },
                "account_origin_id": {
                    "description": "AccountOriginID - identificador da conta de origem.",
                    "type": "string"
                },
                "amount": {
                    "description": "Amount - valor transferido.",
                    "type": "integer"
                },
                "created_at": {
                    "description": "CreatedAt - data de criação da transferência.",
                    "type": "string"
                },
                "id": {
                    "description": "ID - identificador único da transferência.",
                    "type": "string"
                }
            }
        },
        "handler.CreateAccountRequest": {
            "type": "object",
            "required": [
                "balance",
                "cpf",
                "name",
                "password"
            ],
            "properties": {
                "balance": {
                    "description": "Balance - saldo inicial da conta.",
                    "type": "integer"
                },
                "cpf": {
                    "description": "CPF - número do CPF do titular da conta.",
                    "type": "string"
                },
                "name": {
                    "description": "Name - nome do titular da conta.",
                    "type": "string"
                },
                "password": {
                    "description": "Password - senha de acesso a conta.",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Challenge Accounts API",
	Description:      "Implementação de API para o desafio de backend.\nA API é responsável por gerenciar contas e transferências\nentre contas.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
