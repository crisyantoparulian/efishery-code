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
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/display/resource-aggregation": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "This api is for display aggregation of resources",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Resources"
                ],
                "summary": "Resources Aggregation",
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
                                            "$ref": "#/definitions/fetchresourceaggregation.InportResponse"
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
        },
        "/api/v1/display/resource-with-price": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "This api is for display resources with price USD",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Resources"
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
        },
        "/api/v1/user": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "This api is for display resources with price USD",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Display user claim",
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
                                            "$ref": "#/definitions/displayoneuser.InportResponse"
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
        "displayoneuser.InportResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "type": "object",
                    "additionalProperties": true
                }
            }
        },
        "fetchresourceaggregation.Aggregate": {
            "type": "object",
            "properties": {
                "avg": {
                    "type": "number"
                },
                "max": {
                    "type": "integer"
                },
                "med": {
                    "type": "number"
                },
                "min": {
                    "type": "integer"
                }
            }
        },
        "fetchresourceaggregation.AggregateResources": {
            "type": "object",
            "properties": {
                "aggregate_price": {
                    "$ref": "#/definitions/fetchresourceaggregation.Aggregate"
                },
                "aggregate_size": {
                    "$ref": "#/definitions/fetchresourceaggregation.Aggregate"
                },
                "area_provinsi": {
                    "type": "string"
                },
                "price": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "size": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "week": {
                    "type": "integer"
                }
            }
        },
        "fetchresourceaggregation.InportResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/fetchresourceaggregation.AggregateResources"
                    }
                }
            }
        },
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
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Fetch apps",
	Description:      "This API is for fetch & mapping resources",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
