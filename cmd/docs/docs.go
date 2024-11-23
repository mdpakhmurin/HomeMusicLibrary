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
        "/song": {
            "post": {
                "description": "Добавляет новую песню",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Добавить новую песню",
                "parameters": [
                    {
                        "description": "Данные о песне в формате JSON",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/SongCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Сообщение о успешном добавлении песни",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка при обработке запроса",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/song/": {
            "put": {
                "description": "Добавляет новую песню",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Обновляет песню с заданным названием и группой",
                "parameters": [
                    {
                        "description": "Данные о песне в формате JSON",
                        "name": "song",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/SongUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Сообщение о успешном добавлении песни",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка при обработке запроса",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет песню по указанным параметрам: group и name.",
                "produces": [
                    "application/json"
                ],
                "summary": "Удалить песню",
                "parameters": [
                    {
                        "type": "string",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "song",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Сообщение об успешном удалении песни",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Ошибка при обработке запроса",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/song/info": {
            "get": {
                "description": "Получение информации о песне по заданной группе и названию",
                "produces": [
                    "application/json"
                ],
                "summary": "Получение информации о песне",
                "parameters": [
                    {
                        "type": "string",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "song",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/SongDetail"
                        }
                    }
                }
            }
        },
        "/song/search": {
            "get": {
                "description": "Поиск песни по заданным параметрам",
                "produces": [
                    "application/json"
                ],
                "summary": "Поиск песни",
                "parameters": [
                    {
                        "type": "string",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "link",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "releaseDate",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "song",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "text",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/SongDescription"
                            }
                        }
                    }
                }
            }
        },
        "/song/verses": {
            "get": {
                "description": "Получение куплетов песни по заданным параметрам",
                "produces": [
                    "application/json"
                ],
                "summary": "Получение текста песни",
                "parameters": [
                    {
                        "type": "string",
                        "name": "group",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "song",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "SongCreate": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "description": "wewe",
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "SongDescription": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                }
            }
        },
        "SongDetail": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "SongUpdate": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "releaseDate": {
                    "type": "string"
                },
                "song": {
                    "type": "string"
                },
                "text": {
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
	Title:            "Home Music Library API",
	Description:      "This is a simple API for managing home music library.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}