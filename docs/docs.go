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
        "/account/customer-name": {
            "get": {
                "description": "Get Customer Name by Account Number",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Get Customer Name by Account Number",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Account payload",
                        "name": "accountNumber",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-model_GetCustomerNameByAccountNumberResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    }
                }
            }
        },
        "/account/internal-transfer": {
            "post": {
                "description": "Transfer from internal account to internal account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Transfer",
                "parameters": [
                    {
                        "description": "Account payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.InternalTransferRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    }
                }
            }
        },
        "/auth/forgot-password": {
            "post": {
                "description": "Set a new password after OTP verification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auths"
                ],
                "summary": "Set Password",
                "parameters": [
                    {
                        "description": "Set Password payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SetPasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    }
                }
            }
        },
        "/auth/forgot-password/otp": {
            "post": {
                "description": "Send OTP to user email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auths"
                ],
                "summary": "Send OTP to Mail",
                "parameters": [
                    {
                        "description": "Send OTP payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SendOTPRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    }
                }
            }
        },
        "/auth/forgot-password/verify-otp": {
            "post": {
                "description": "Verify OTP with email and otp",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auths"
                ],
                "summary": "Verify OTP",
                "parameters": [
                    {
                        "description": "Verify OTP payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.VerifyOTPRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login to account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auths"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Auth payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-entity_Customer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register to account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auths"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Auth payload",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    }
                }
            }
        },
        "/core/estimate-transfer-fee": {
            "get": {
                "description": "Estimate the internal transfer fee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cores"
                ],
                "summary": "EstimateTransferFee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Amount to estimate",
                        "name": "amount",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-int64"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpcommon.HttpResponse-any"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Customer": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
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
                "phoneNumber": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "httpcommon.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "httpcommon.HttpResponse-any": {
            "type": "object",
            "properties": {
                "data": {},
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/httpcommon.Error"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "httpcommon.HttpResponse-entity_Customer": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/entity.Customer"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/httpcommon.Error"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "httpcommon.HttpResponse-int64": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "integer"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/httpcommon.Error"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "httpcommon.HttpResponse-model_GetCustomerNameByAccountNumberResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.GetCustomerNameByAccountNumberResponse"
                },
                "errors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/httpcommon.Error"
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "model.GetCustomerNameByAccountNumberResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "model.InternalTransferRequest": {
            "type": "object",
            "required": [
                "amount",
                "isSourceFee",
                "sourceAccountNumber",
                "targetAccountNumber"
            ],
            "properties": {
                "amount": {
                    "type": "integer",
                    "minimum": 0
                },
                "isSourceFee": {
                    "type": "boolean"
                },
                "sourceAccountNumber": {
                    "type": "string"
                },
                "targetAccountNumber": {
                    "type": "string"
                }
            }
        },
        "model.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "recaptchaToken"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 10
                },
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 8
                },
                "recaptchaToken": {
                    "type": "string"
                }
            }
        },
        "model.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "phoneNumber"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 10
                },
                "name": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 5
                },
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 8
                },
                "phoneNumber": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 10
                }
            }
        },
        "model.SendOTPRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 10
                }
            }
        },
        "model.SetPasswordRequest": {
            "type": "object",
            "required": [
                "email",
                "otp",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 10
                },
                "otp": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6
                },
                "password": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 8
                }
            }
        },
        "model.VerifyOTPRequest": {
            "type": "object",
            "required": [
                "email",
                "otp"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 10
                },
                "otp": {
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6
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
