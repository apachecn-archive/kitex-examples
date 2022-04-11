// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
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
        "contact": {
            "name": "Bodhisatan",
            "url": "https://github.com/bodhisatan",
            "email": "bodhisatanyao@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/product/brand_create": {
            "post": {
                "description": "商家绑定品牌，返回品牌ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块-品牌子模块"
                ],
                "summary": "商家绑定品牌",
                "parameters": [
                    {
                        "description": "品牌信息",
                        "name": "shopSettleParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.BrandAddParam"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer $token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        },
        "/product/brand_del": {
            "post": {
                "description": "商家删除品牌",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块-品牌子模块"
                ],
                "summary": "商家删除品牌",
                "parameters": [
                    {
                        "description": "品牌信息",
                        "name": "shopSettleParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.BrandDelParam"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer $token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        },
        "/product/brand_update": {
            "post": {
                "description": "商家更新品牌信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块-品牌子模块"
                ],
                "summary": "商家更新品牌信息",
                "parameters": [
                    {
                        "description": "品牌信息",
                        "name": "shopSettleParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.BrandEditParam"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer $token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        },
        "/product/get_brands": {
            "get": {
                "description": "商家绑定品牌查询",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品模块-品牌子模块"
                ],
                "summary": "商家绑定品牌查询",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer $token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        },
        "/shop/id": {
            "get": {
                "description": "通过用户ID查询商家ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商家模块"
                ],
                "summary": "商家ID查询",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer $token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        },
        "/shop/settle": {
            "post": {
                "description": "商家入驻，返回ShopID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商家模块"
                ],
                "summary": "商家入驻",
                "parameters": [
                    {
                        "description": "入驻材料",
                        "name": "shopSettleParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.ShopSettleParam"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer $token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "账号信息",
                        "name": "userParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UserParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.LoginResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户模块"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "注册信息",
                        "name": "userParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.UserParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.BrandAddParam": {
            "type": "object",
            "properties": {
                "brand_name": {
                    "type": "string"
                },
                "brand_story": {
                    "type": "string"
                },
                "logo": {
                    "type": "string"
                }
            }
        },
        "handlers.BrandDelParam": {
            "type": "object",
            "required": [
                "brand_id"
            ],
            "properties": {
                "brand_id": {
                    "type": "integer"
                }
            }
        },
        "handlers.BrandEditParam": {
            "type": "object",
            "required": [
                "brand_id"
            ],
            "properties": {
                "brand_id": {
                    "type": "integer"
                },
                "brand_name": {
                    "type": "string"
                },
                "brand_story": {
                    "type": "string"
                },
                "logo": {
                    "type": "string"
                }
            }
        },
        "handlers.LoginResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "expire": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "handlers.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "handlers.ShopSettleParam": {
            "type": "object",
            "properties": {
                "shop_name": {
                    "type": "string"
                }
            }
        },
        "handlers.UserParam": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
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
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Mall",
	Description:      "This is a mall demo using Gin and KiteX.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
