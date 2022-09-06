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
        "/api/v1/demo/balance": {
            "get": {
                "tags": [
                    "Demo"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/demo/name": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Demo"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/toekn/refresh": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "parameters": [
                    {
                        "description": "refresh token",
                        "name": "refreshToken",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RefershRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jwt_auth.AuthToken"
                        }
                    }
                }
            }
        },
        "/api/v1/toekn/revoke": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jwt_auth.AuthToken"
                        }
                    }
                }
            }
        },
        "/api/v1/users/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "description": "Add account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jwt_auth.AuthToken"
                        }
                    }
                }
            }
        },
        "/api/v1/users/info": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Member"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "description": "update user info",
                        "name": "userInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/api/v1/users/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "parameters": [
                    {
                        "description": "Add account",
                        "name": "userInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/jwt_auth.AuthToken"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.LoginRequest": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string",
                    "example": "Jack"
                },
                "password": {
                    "type": "string",
                    "example": "12345"
                }
            }
        },
        "controller.RefershRequest": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "xxcjdjfcidasjcodioi"
                }
            }
        },
        "jwt_auth.AuthToken": {
            "type": "object",
            "properties": {
                "accessExp": {
                    "type": "integer",
                    "example": 600
                },
                "accessToken": {
                    "type": "string",
                    "example": "dkdke3klwlwkkf..."
                },
                "refreshExp": {
                    "type": "integer",
                    "example": 86400
                },
                "refreshToken": {
                    "type": "string",
                    "example": "dkdke3klwlwkkf..."
                },
                "tokenType": {
                    "type": "string",
                    "example": "Bearer"
                }
            }
        },
        "model.Member": {
            "type": "object",
            "properties": {
                "account": {
                    "type": "string"
                },
                "address": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isVerify": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "model.UpdateUser": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "地球"
                },
                "gender": {
                    "type": "string",
                    "example": "m"
                },
                "name": {
                    "type": "string",
                    "example": "Jack"
                },
                "phone": {
                    "type": "string",
                    "example": "0987654321"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8088",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Gin Swagger Demo",
	Description:      "Swagger API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
