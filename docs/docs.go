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
            "name": "Minh Thuong"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/tags": {
            "get": {
                "description": "Retrieve a list of all tags with pagination support",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Retrieve all tags",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number (default is 1)",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of items per page (default is 10)",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Paginated list of tags",
                        "schema": {
                            "$ref": "#/definitions/tagresponse.FindAllTagResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new tag based on the provided JSON data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Create a new tag",
                "parameters": [
                    {
                        "description": "Tag data",
                        "name": "tag",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/tagrequest.CreateTagRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Tag created successfully",
                        "schema": {
                            "$ref": "#/definitions/model.Tag"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Tag": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "tag name"
                }
            }
        },
        "tagrequest.CreateTagRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "tagresponse.FindAllTagResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Tag"
                    }
                },
                "page": {
                    "type": "integer"
                },
                "totalItems": {
                    "type": "integer"
                },
                "totalPages": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Cake ATS API",
	Description:      "This is a sample server Cake ATS API server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
