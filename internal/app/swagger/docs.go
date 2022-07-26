// Package swagger GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package swagger

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Eason Chai",
            "email": "i@hackerchai.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/threats": {
            "get": {
                "description": "获取所有的威胁IP信息, 并分页的形式呈现",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Threat"
                ],
                "summary": "获取所有的威胁IP信息",
                "operationId": "get-threats",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "当前查询页码(默认 1)",
                        "name": "current",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "当前查询每页显示条数(默认 10)",
                        "name": "page_sizet",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.ThreatsResponse"
                        }
                    },
                    "400": {
                        "description": "{code:1002,status_code:401,message:no permission}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "{code:1001,status_code:504,message:internal server error}",
                        "schema": {
                            "$ref": "#/definitions/schema.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ent.Threat": {
            "type": "object",
            "properties": {
                "ctime": {
                    "description": "Ctime holds the value of the \"ctime\" field.",
                    "type": "integer"
                },
                "domain_count": {
                    "description": "DomainCount holds the value of the \"domain_count\" field.",
                    "type": "integer"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                },
                "ip": {
                    "description": "IP holds the value of the \"ip\" field.",
                    "type": "string"
                },
                "itel_count": {
                    "description": "ItelCount holds the value of the \"itel_count\" field.",
                    "type": "integer"
                },
                "judge": {
                    "description": "Judge holds the value of the \"judge\" field.",
                    "type": "integer"
                },
                "poc": {
                    "description": "Poc holds the value of the \"poc\" field.",
                    "type": "boolean"
                },
                "source": {
                    "description": "Source holds the value of the \"source\" field.",
                    "type": "integer"
                },
                "tag_count": {
                    "description": "TagCount holds the value of the \"tag_count\" field.",
                    "type": "integer"
                },
                "threat_id_info": {
                    "description": "ThreatIDInfo holds the value of the \"threat_id_info\" field.",
                    "type": "string"
                }
            }
        },
        "schema.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "schema.PaginationResponse": {
            "type": "object",
            "properties": {
                "current": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "schema.ThreatsPage": {
            "type": "object",
            "properties": {
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Threat"
                    }
                },
                "pagnition": {
                    "$ref": "#/definitions/schema.PaginationResponse"
                }
            }
        },
        "schema.ThreatsResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/schema.ThreatsPage"
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "threatbook-ip-web",
	Description: "Web for craling threatbook community to obtain threat ip from feeds.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
