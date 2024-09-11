// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/users/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "Data request",
                        "name": "request-data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/usermodel.UserCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Verify id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/verify-registration": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Verify registration",
                "parameters": [
                    {
                        "description": "Verify request data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/verifier.VerifyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error details",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "usermodel.UserCreateReq": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "first_name": {
                    "type": "string",
                    "example": "Van"
                },
                "last_name": {
                    "type": "string",
                    "example": "Nguyen"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "example": "0914590038"
                },
                "verify_by": {
                    "description": "confirm by sms or email",
                    "type": "string",
                    "example": "sms"
                }
            }
        },
        "verifier.VerifyRequest": {
            "type": "object",
            "required": [
                "verify_code",
                "verify_id",
                "verify_info"
            ],
            "properties": {
                "verify_code": {
                    "type": "string",
                    "example": "1234"
                },
                "verify_id": {
                    "type": "string",
                    "example": "abc31233"
                },
                "verify_info": {
                    "type": "string",
                    "example": "email or phone number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
