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
        "/v1/admin/approve/{transactionid}": {
            "post": {
                "description": "Approve Transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Approve Transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "transaction_id",
                        "name": "transactionid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/admin/customer": {
            "get": {
                "description": "Manage Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Manage Customer",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.ManageCustomer"
                        }
                    }
                }
            },
            "put": {
                "description": "Update customer data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Update customer data",
                "parameters": [
                    {
                        "description": "Customer",
                        "name": "Customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.ManageCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.ManageCustomer"
                        }
                    }
                }
            }
        },
        "/v1/admin/customerpoint": {
            "get": {
                "description": "Manage Customer Point",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Manage Customer Point",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.UserPoin"
                        }
                    }
                }
            },
            "put": {
                "description": "Update customer point",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Update customer point",
                "parameters": [
                    {
                        "description": "Customer Point",
                        "name": "UpdateCustomerPoint",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.UpdateCustomerPoint"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/admin/dashboard": {
            "get": {
                "description": "Dashboard Admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Dashboard Admin",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.Dashboard"
                        }
                    }
                }
            }
        },
        "/v1/admin/historycustomer": {
            "get": {
                "description": "Customer history",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Customer history",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.CustomerHistory"
                        }
                    }
                }
            }
        },
        "/v1/admin/login": {
            "post": {
                "description": "Login Admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Admin",
                        "name": "LoginAdmin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.LoginAdmin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/admin.LoginAdmin"
                        }
                    }
                }
            }
        },
        "/v1/customer": {
            "put": {
                "description": "Update Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Update Customer",
                "parameters": [
                    {
                        "description": "UpdateCustomer",
                        "name": "Customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dummy.UpdateCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dummy.UpdateCustomer"
                        }
                    }
                }
            }
        },
        "/v1/detailhistory/{transactionid}": {
            "get": {
                "description": "History/transaction Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Detail History/transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id detail history",
                        "name": "transactionid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dummy.DetailTransactionCustomer"
                        }
                    }
                }
            }
        },
        "/v1/emoney": {
            "post": {
                "description": "Redeem Emoney/Cashout",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RedeemPoint"
                ],
                "summary": "Redeem Emoney/Cashout",
                "parameters": [
                    {
                        "description": "RedeemEmoney",
                        "name": "RedeemEmoney",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dummy.RedeemEmoney"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dummy.RedeemEmoney"
                        }
                    }
                }
            }
        },
        "/v1/history/{idcustomer}": {
            "get": {
                "description": "History Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
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
                        "description": "id customer",
                        "name": "idcustomer",
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
                "description": "Login Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Customer",
                        "name": "Customer",
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
                            "$ref": "#/definitions/dummy.Customer"
                        }
                    }
                }
            }
        },
        "/v1/mitra/history": {
            "get": {
                "description": "History Mitra",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mitra"
                ],
                "summary": "History Mitra",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/mitra/inputpoin": {
            "post": {
                "description": "Input Poin Mitra",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mitra"
                ],
                "summary": "Input Poin Mitra",
                "parameters": [
                    {
                        "description": "InputPoinMitra",
                        "name": "InputPoin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mitra.InputPoinMitra"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/v1/mitra/login": {
            "post": {
                "description": "Login Mitra",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mitra"
                ],
                "summary": "Login Mitra",
                "parameters": [
                    {
                        "description": "mitra",
                        "name": "mitra",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mitra.AuthMitra"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/mitra.Mitra"
                        }
                    }
                }
            }
        },
        "/v1/mitra/register": {
            "post": {
                "description": "Register Mitra",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mitra"
                ],
                "summary": "Register Mitra",
                "parameters": [
                    {
                        "description": "MitraRegister",
                        "name": "RegisterMitra",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/mitra.MitraRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/mitra.MitraRegister"
                        }
                    }
                }
            }
        },
        "/v1/paketdata": {
            "post": {
                "description": "Redeem PaketData",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RedeemPoint"
                ],
                "summary": "Redeem PaketData",
                "parameters": [
                    {
                        "description": "RedeemPaketData",
                        "name": "RedeemPaketData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dummy.RedeemPulsaPaketData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dummy.RedeemPulsaPaketData"
                        }
                    }
                }
            }
        },
        "/v1/pulsa": {
            "post": {
                "description": "Redeem Pulsa",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "RedeemPoint"
                ],
                "summary": "Redeem Pulsa",
                "parameters": [
                    {
                        "description": "RedeemPulsa",
                        "name": "RedeemPulsa",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dummy.RedeemPulsaPaketData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dummy.RedeemPulsaPaketData"
                        }
                    }
                }
            }
        },
        "/v1/register": {
            "post": {
                "description": "Register Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Register",
                        "name": "RegisterCustomer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dummy.Register"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dummy.Register"
                        }
                    }
                }
            }
        },
        "/v1/stockproduct": {
            "get": {
                "description": "Get Stock Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get Stock Product",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseGetProduct"
                        }
                    }
                }
            }
        },
        "/v1/stockproduct/{id}": {
            "put": {
                "description": "Update Stock Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update Stock Product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id product",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Stock Product",
                        "name": "Product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dummy.InputStockProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.ResponseGetProduct"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "admin.CustomerHistory": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "keterangan": {
                    "type": "string"
                },
                "nomor": {
                    "type": "string"
                },
                "poin_redeem": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "tanggal": {
                    "type": "string"
                }
            }
        },
        "admin.Dashboard": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "keterangan": {
                    "type": "string"
                },
                "nomor": {
                    "type": "string"
                },
                "status_transaction": {
                    "type": "string"
                },
                "transaction_id": {
                    "type": "string"
                }
            }
        },
        "admin.LoginAdmin": {
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
        "admin.ManageCustomer": {
            "type": "object",
            "required": [
                "email",
                "name",
                "no_hp",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "no_hp": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "admin.UpdateCustomerPoint": {
            "type": "object",
            "required": [
                "poin_account"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "poin_account": {
                    "type": "integer"
                }
            }
        },
        "admin.UserPoin": {
            "type": "object",
            "properties": {
                "alamat": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "poin": {
                    "type": "integer"
                }
            }
        },
        "dummy.AuthLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dummy.Customer": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "no_hp": {
                    "type": "string"
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
        "dummy.DetailTransactionCustomer": {
            "type": "object",
            "properties": {
                "bank_provider": {
                    "type": "string"
                },
                "createdat": {
                    "type": "string"
                },
                "jenis_transaction": {
                    "type": "string"
                },
                "nomor": {
                    "type": "string"
                },
                "poin_account": {
                    "type": "integer"
                },
                "poin_redeem": {
                    "type": "integer"
                },
                "transaction_id": {
                    "type": "string"
                }
            }
        },
        "dummy.History": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "keterangan": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "tanggal": {
                    "type": "string"
                },
                "transaction_id": {
                    "type": "string"
                }
            }
        },
        "dummy.InputStockProduct": {
            "type": "object",
            "properties": {
                "stock": {
                    "type": "integer"
                }
            }
        },
        "dummy.RedeemEmoney": {
            "type": "object",
            "required": [
                "amount",
                "bank_provider",
                "customer_id",
                "nomor",
                "poin_account",
                "poin_redeem"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "an_rekening": {
                    "type": "string"
                },
                "bank_provider": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "nomor": {
                    "type": "string"
                },
                "poin_account": {
                    "type": "integer"
                },
                "poin_redeem": {
                    "type": "integer"
                }
            }
        },
        "dummy.RedeemPulsaPaketData": {
            "type": "object",
            "required": [
                "amount",
                "bank_provider",
                "customer_id",
                "nomor",
                "poin_account",
                "poin_redeem"
            ],
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "bank_provider": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "integer"
                },
                "nomor": {
                    "type": "string"
                },
                "poin_account": {
                    "type": "integer"
                },
                "poin_redeem": {
                    "type": "integer"
                }
            }
        },
        "dummy.Register": {
            "type": "object",
            "required": [
                "email",
                "name",
                "no_hp",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "no_hp": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dummy.StockProduct": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "product": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "dummy.UpdateCustomer": {
            "type": "object",
            "required": [
                "name",
                "no_hp",
                "password"
            ],
            "properties": {
                "idcustomer": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "no_hp": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "mitra.AuthMitra": {
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
        "mitra.InputPoinMitra": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "idcustomer": {
                    "type": "integer"
                },
                "idmitra": {
                    "type": "integer"
                }
            }
        },
        "mitra.Mitra": {
            "type": "object",
            "properties": {
                "alamat": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nama_toko": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "mitra.MitraRegister": {
            "type": "object",
            "properties": {
                "alamat": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nama_toko": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "response.ResponseGetProduct": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dummy.StockProduct"
                    }
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
