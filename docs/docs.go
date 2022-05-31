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
                "responses": {}
            }
        },
        "/v1/product/cashout": {
            "get": {
                "description": "Product CashOut user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Product CashOut",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dummy.ProductCashout"
                            }
                        }
                    }
                }
            }
        },
        "/v1/product/emoney": {
            "get": {
                "description": "Product EMoney user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Product EMoney",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dummy.ProductEmoney"
                            }
                        }
                    }
                }
            }
        },
        "/v1/product/paketdata": {
            "get": {
                "description": "Product PaketData user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Product PaketData",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dummy.ProductPaketData"
                            }
                        }
                    }
                }
            }
        },
        "/v1/product/pulsa": {
            "get": {
                "description": "Product Pulsa user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Product Pulsa",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dummy.ProductPulsa"
                            }
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
                "responses": {}
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
        },
        "dummy.ProductCashout": {
            "type": "object",
            "properties": {
                "harga": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "poin": {
                    "type": "integer"
                }
            }
        },
        "dummy.ProductEmoney": {
            "type": "object",
            "properties": {
                "harga": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "poin": {
                    "type": "integer"
                }
            }
        },
        "dummy.ProductPaketData": {
            "type": "object",
            "properties": {
                "Internet": {
                    "type": "string"
                },
                "harga": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "kuota": {
                    "type": "string"
                },
                "poin": {
                    "type": "integer"
                }
            }
        },
        "dummy.ProductPulsa": {
            "type": "object",
            "properties": {
                "harga": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "poin": {
                    "type": "integer"
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
