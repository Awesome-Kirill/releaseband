// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "email": "k.a.stulnikov@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/game/{id}/calculate": {
            "get": {
                "description": "Return game result",
                "produces": [
                    "application/json"
                ],
                "summary": "Return game result",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Game ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Result"
                        }
                    }
                }
            }
        },
        "/game/{id}/lines": {
            "post": {
                "description": "Return game result",
                "produces": [
                    "application/json"
                ],
                "summary": "Return game result",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Game ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Some ID",
                        "name": "some_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.WinLine"
                            }
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/game/{id}/payouts": {
            "post": {
                "description": "Return game result",
                "produces": [
                    "application/json"
                ],
                "summary": "Return game result",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Game ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Some ID",
                        "name": "some_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.PayoutSymbol"
                            }
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/game/{id}/reels": {
            "post": {
                "description": "Return game result",
                "produces": [
                    "application/json"
                ],
                "summary": "Return game result",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Game ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Some ID",
                        "name": "some_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "array",
                                "items": {
                                    "type": "string"
                                }
                            }
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "domain.Line": {
            "type": "object",
            "properties": {
                "line": {
                    "type": "integer"
                },
                "payout": {
                    "type": "integer"
                }
            }
        },
        "domain.PayoutSymbol": {
            "type": "object",
            "properties": {
                "payout": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "domain.Position": {
            "type": "object",
            "properties": {
                "col": {
                    "type": "integer"
                },
                "row": {
                    "type": "integer"
                }
            }
        },
        "domain.Result": {
            "type": "object",
            "properties": {
                "lines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Line"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "domain.WinLine": {
            "type": "object",
            "properties": {
                "line": {
                    "type": "integer"
                },
                "positions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Position"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server Petstore server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}