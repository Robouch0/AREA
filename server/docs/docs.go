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
                    "About"
                ],
                "summary": "List of handled services",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.AboutInfos"
                        }
                    }
                }
            }
        },
        "/area/activate": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
        "/area/create/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
        "/area/create/{service}": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
        "/area/list": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Pong",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Ping"
                ],
                "summary": "Prints pong",
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
        "/token": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all the tokens of the current logged user",
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
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.TokenInformations"
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
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                "parameters": [
                    {
                        "description": "Token creation request informations",
                        "name": "tokenCreateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.TokenCreateRequest"
                        }
                    }
                ],
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
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
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
                "parameters": [
                    {
                        "type": "string",
                        "description": "Remote Service Name",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    }
                ],
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
        "/token/{provider}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a the token associated to the remote provider of the user",
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
                        "description": "Remote Service Name",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    }
                ],
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
        "/user/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new user in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User's information",
                        "name": "userInfos",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserInformations"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.UserInformations"
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
        "/user/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get user's information based on his ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User By ID",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.UserInformations"
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
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update some informations about the user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User datas",
                "parameters": [
                    {
                        "description": "Updatable user's informations",
                        "name": "updatableDatas",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdatableUserData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UpdatableUserData"
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
        "controllers.AboutInfos": {
            "type": "object",
            "properties": {
                "client": {
                    "$ref": "#/definitions/controllers.clientInfos"
                },
                "server": {
                    "$ref": "#/definitions/controllers.serverInfos"
                }
            }
        },
        "controllers.TokenCreateRequest": {
            "type": "object",
            "properties": {
                "provider": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
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
        "controllers.UserInformations": {
            "type": "object",
            "properties": {
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
                }
            }
        },
        "controllers.clientInfos": {
            "type": "object",
            "properties": {
                "host": {
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
        "controllers.microservice": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.serverInfos": {
            "type": "object",
            "properties": {
                "current_time": {
                    "type": "integer"
                },
                "services": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.serverService"
                    }
                }
            }
        },
        "controllers.serverService": {
            "type": "object",
            "properties": {
                "actions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.microservice"
                    }
                },
                "name": {
                    "type": "string"
                },
                "reactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.microservice"
                    }
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
        "models.UpdatableUserData": {
            "type": "object",
            "properties": {
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
                    "description": "Type of service action or reaction",
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
