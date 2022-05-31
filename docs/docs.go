// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/detailhistory/{id}": {
            "get": {
                "description": "History/transaction User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Detail History/transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id detail history",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dummy.DetailTransaction"
                        }
                    }
                }
            }
        },
        "/v1/history/{iduser}": {
            "get": {
                "description": "History User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "History",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "id user",
                        "name": "iduser",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dummy.History"
                        }
                    }
                }
            }
        },
        "/v1/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dummy.AuthLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dummy.Login"
                        }
                    }
                }
            }
        },
        "/v1/order/cashout": {
            "post": {
                "description": "Register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserOrder"
                ],
                "summary": "Order Cashout",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/order/emoney": {
            "post": {
                "description": "Emoney user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserOrder"
                ],
                "summary": "Order Emoney",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/order/paketdata": {
            "post": {
                "description": "PaketData user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserOrder"
                ],
                "summary": "Order PaketData",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/order/pulsa": {
            "post": {
                "description": "Pulsa user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserOrder"
                ],
                "summary": "Order Pulsa",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/v1/register": {
            "post": {
                "description": "Register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dummy.AuthLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dummy.DetailTransaction": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "jenis_transaction": {
                    "type": "string"
                },
                "nama_bank": {
                    "type": "string"
                },
                "no_rekening": {
                    "type": "integer"
                },
                "poin_account": {
                    "type": "integer"
                },
                "poin_redeem": {
                    "type": "integer"
                }
            }
        },
        "dummy.History": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "tanggal": {
                    "type": "string"
                },
                "tipe_transaksi": {
                    "type": "string"
                }
            }
        },
        "dummy.Login": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "pin": {
                    "type": "integer"
                },
                "poin": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "api-dummy.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API-Loyalty Point Agent",
	Description:      "Berikut API-Loyalty Point Agent",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
