// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/v1/display-with-price": {
            "get": {
                "description": "This api is for display resources with price USD",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Resources With Price USD",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/restapi.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/fetchresourcewithprice.InportResponse"
                                        },
                                        "errorCode": {
                                            "type": "string"
                                        },
                                        "errorMessage": {
                                            "type": "string"
                                        },
                                        "success": {
                                            "type": "boolean"
                                        },
                                        "traceId": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "fetchresourcewithprice.Area": {
            "type": "object",
            "properties": {
                "kota": {
                    "type": "string"
                },
                "provinsi": {
                    "type": "string"
                }
            }
        },
        "fetchresourcewithprice.InportResponse": {
            "type": "object",
            "properties": {
                "resources": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/fetchresourcewithprice.ResourceWithPrice"
                    }
                }
            }
        },
        "fetchresourcewithprice.Price": {
            "type": "object",
            "properties": {
                "idr": {
                    "type": "string"
                },
                "usd": {
                    "type": "string"
                }
            }
        },
        "fetchresourcewithprice.ResourceWithPrice": {
            "type": "object",
            "properties": {
                "area": {
                    "$ref": "#/definitions/fetchresourcewithprice.Area"
                },
                "komoditas": {
                    "type": "string"
                },
                "price": {
                    "$ref": "#/definitions/fetchresourcewithprice.Price"
                },
                "size": {
                    "type": "string"
                },
                "tgl_parsed": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "restapi.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errorCode": {
                    "type": "string"
                },
                "errorMessage": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "traceId": {
                    "type": "string"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}