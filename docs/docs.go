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
        "/gps": {
            "post": {
                "description": "Create\tgps",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gps"
                ],
                "summary": "Create\tgps",
                "parameters": [
                    {
                        "description": "gps request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateGPSInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.Error"
                        }
                    }
                }
            }
        },
        "/gyroscope": {
            "post": {
                "description": "Create\tgyroscope",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gyroscopes"
                ],
                "summary": "Create\tgyroscope",
                "parameters": [
                    {
                        "description": "gyroscope request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateGyroscopeInputDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/web.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/web.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateGPSInputDTO": {
            "type": "object",
            "required": [
                "latitude",
                "longitude",
                "mac_address"
            ],
            "properties": {
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "mac_address": {
                    "type": "string"
                }
            }
        },
        "dto.CreateGyroscopeInputDTO": {
            "type": "object",
            "required": [
                "mac_address",
                "model",
                "name",
                "x",
                "y",
                "z"
            ],
            "properties": {
                "mac_address": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "x": {
                    "type": "number"
                },
                "y": {
                    "type": "number"
                },
                "z": {
                    "type": "number"
                }
            }
        },
        "web.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1/telemetry/",
	Schemes:          []string{},
	Title:            "telemetry Service API",
	Description:      "This is a telemetry microservice API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
