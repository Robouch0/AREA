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
        "/about.json": {
            "get": {
                "description": "json giving the list of handled action-reaction services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Utils"
                ],
                "summary": "List of handled services",
                "responses": {}
            }
        },
        "/area/activate": {
            "put": {
                "description": "Activate/Deactivate user's area",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area"
                ],
                "summary": "Activate/Deactivate an area",
                "parameters": [
                    {
                        "description": "Informations about the activation of an area",
                        "name": "area",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/areas.areaActivateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serviceinterface.SetActivatedResponseStatus"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/area/list": {
            "get": {
                "description": "List all user's area",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area"
                ],
                "summary": "List User's area",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/areas.userArea"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/create/list": {
            "get": {
                "description": "List all available areas",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area"
                ],
                "summary": "List available areas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/serviceinterface.ServiceStatus"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/create/{service}": {
            "get": {
                "description": "Register a new Area in the application",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area"
                ],
                "summary": "Create a new Area",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "service",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Full body of an Area Scenario",
                        "name": "area",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AreaScenario"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serviceinterface.ActionResponseStatus"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/login/": {
            "post": {
                "description": "Login a user if he has the correct credentials and returns the tokens and the user_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Sign-In",
                "parameters": [
                    {
                        "description": "Credentials of the user who wants to connect",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.credentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/log_types.UserLogInfos"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/oauth/": {
            "post": {
                "description": "Create account with code from redirect url",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Create account with oauth",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/log_types.UserLogInfos"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/oauth/{service}": {
            "get": {
                "description": "Get the oauth redirect url for a service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Get an oauth url for a service",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Redirect URL for the oauth",
                        "name": "redirect_uri",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the service to use oauth with",
                        "name": "service",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "pong",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "prints pong",
                "responses": {
                    "200": {
                        "description": "pong"
                    }
                }
            }
        },
        "/sign-up/": {
            "post": {
                "description": "register an account by giving credentials",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "summary": "Sign-up a new account",
                "parameters": [
                    {
                        "description": "New User informations to sign-up to the app",
                        "name": "newUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.userSignUp"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.userSignUp"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/token/": {
            "post": {
                "description": "Get the tokens from a user_id and a provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Get user's token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the user",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Provider of the Remote Service",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "delete": {
                "description": "Delete a token from a user_id and a provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Delete a token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.TokenInformations"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/token/create/": {
            "post": {
                "description": "Create a token from a user_id and a provider",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Create a token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Token"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/token/{user_id}": {
            "get": {
                "description": "Get all the tokens from a user_id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Get all the tokens from a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id of the user",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Token"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/webhook/{service}/{microservice}/{action_id}": {
            "post": {
                "description": "Webhook Endpoint for the remote services payloads",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Area"
                ],
                "summary": "Webhook Endpoint",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Service Name",
                        "name": "service",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Microservice Name",
                        "name": "microservice",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Action ID for the reaction service",
                        "name": "action_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "areas.areaActivateRequest": {
            "type": "object",
            "properties": {
                "activated": {
                    "type": "boolean"
                },
                "area_id": {
                    "type": "integer"
                }
            }
        },
        "areas.userArea": {
            "type": "object",
            "properties": {
                "action": {
                    "$ref": "#/definitions/serviceinterface.ServiceStatus"
                },
                "activated": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "reactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serviceinterface.ServiceStatus"
                    }
                }
            }
        },
        "controllers.TokenInformations": {
            "type": "object",
            "properties": {
                "provider": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "controllers.credentials": {
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
        "controllers.userSignUp": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "log_types.UserLogInfos": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Action": {
            "type": "object",
            "properties": {
                "ingredients": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "microservice": {
                    "type": "string"
                },
                "service": {
                    "type": "string"
                }
            }
        },
        "models.Actions": {
            "type": "object",
            "properties": {
                "action": {
                    "$ref": "#/definitions/models.Action"
                },
                "area_id": {
                    "description": "No anotation here !",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.Area": {
            "type": "object",
            "properties": {
                "action": {
                    "$ref": "#/definitions/models.Actions"
                },
                "activated": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "one_shot": {
                    "type": "boolean"
                },
                "reactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Reactions"
                    }
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.AreaScenario": {
            "type": "object",
            "properties": {
                "action": {
                    "$ref": "#/definitions/models.Action"
                },
                "reaction": {
                    "$ref": "#/definitions/models.Reaction"
                }
            }
        },
        "models.Reaction": {
            "type": "object",
            "properties": {
                "ingredients": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "microservice": {
                    "type": "string"
                },
                "service": {
                    "type": "string"
                }
            }
        },
        "models.Reactions": {
            "type": "object",
            "properties": {
                "area_id": {
                    "description": "No anotation here !",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "prev_out": {
                    "type": "object",
                    "additionalProperties": true
                },
                "reaction": {
                    "$ref": "#/definitions/models.Reaction"
                }
            }
        },
        "models.Token": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "provider": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "areas": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Area"
                    }
                },
                "created_at": {
                    "description": "Useful for log and security purposes",
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "tokens": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Token"
                    }
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "serviceinterface.ActionResponseStatus": {
            "type": "object",
            "properties": {
                "action_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                }
            }
        },
        "serviceinterface.IngredientDescriptor": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "required": {
                    "type": "boolean"
                },
                "type": {
                    "type": "string"
                },
                "value": {}
            }
        },
        "serviceinterface.Ingredients": {
            "type": "object",
            "additionalProperties": {
                "$ref": "#/definitions/serviceinterface.IngredientDescriptor"
            }
        },
        "serviceinterface.MicroserviceDescriptor": {
            "type": "object",
            "properties": {
                "ingredients": {
                    "$ref": "#/definitions/serviceinterface.Ingredients"
                },
                "name": {
                    "description": "Name of the microservice",
                    "type": "string"
                },
                "ref_name": {
                    "description": "Reference Name of the microservice as it is named in the server",
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "serviceinterface.ServiceStatus": {
            "type": "object",
            "properties": {
                "microservices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serviceinterface.MicroserviceDescriptor"
                    }
                },
                "name": {
                    "description": "Name of the service",
                    "type": "string"
                },
                "ref_name": {
                    "description": "Reference Name of the service as it is named in the server",
                    "type": "string"
                }
            }
        },
        "serviceinterface.SetActivatedResponseStatus": {
            "type": "object",
            "properties": {
                "action_id": {
                    "type": "integer"
                },
                "activated": {
                    "type": "boolean"
                },
                "description": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger AREA API",
	Description:      "This is a the document of the Backend routes of the application AREA",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
