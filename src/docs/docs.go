// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/personel/getOrganizationFromPersonel": {
            "get": {
                "description": "GetOrganizationFromPersonel",
                "consumes": [
                    "application/json"
                ],
                "summary": "GetOrganizationFromPersonel",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Models.OrganizasyonSrvcResult"
                        }
                    }
                }
            }
        },
        "/personel/getPersonels": {
            "get": {
                "description": "Get All Personels",
                "consumes": [
                    "application/json"
                ],
                "summary": "Get Personels",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Models.PersonelSrvcResult"
                        }
                    }
                }
            }
        },
        "/student/getOrganizationFromStudent": {
            "get": {
                "description": "GetOrganizationFromStudent",
                "consumes": [
                    "application/json"
                ],
                "summary": "GetOrganizationFromStudent",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Models.OrganizasyonSrvcResult"
                        }
                    }
                }
            }
        },
        "/student/getStudents": {
            "get": {
                "description": "Get All Students",
                "consumes": [
                    "application/json"
                ],
                "summary": "Get Students",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Models.OgrenciSrvcResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Models.Ogrenci": {
            "type": "object",
            "properties": {
                "ad": {
                    "type": "string"
                },
                "birimID": {
                    "type": "string"
                },
                "cepTel": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "ogrenciNo": {
                    "type": "string"
                },
                "ogretimDuzeyi": {
                    "type": "string"
                },
                "ogretimDuzeyiAciklama": {
                    "type": "string"
                },
                "programID": {
                    "type": "string"
                },
                "soyad": {
                    "type": "string"
                },
                "ustBirimID": {
                    "type": "string"
                }
            }
        },
        "Models.OgrenciSrvcResult": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Models.Ogrenci"
                    }
                },
                "message": {
                    "type": "string"
                },
                "result": {
                    "type": "boolean"
                }
            }
        },
        "Models.Organizasyon": {
            "type": "object",
            "properties": {
                "ad": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        },
        "Models.OrganizasyonSrvcResult": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Models.Organizasyon"
                    }
                },
                "message": {
                    "type": "string"
                },
                "result": {
                    "type": "boolean"
                }
            }
        },
        "Models.Personel": {
            "type": "object",
            "properties": {
                "adi": {
                    "type": "string"
                },
                "aktif": {
                    "type": "string"
                },
                "cinsiyet": {
                    "type": "string"
                },
                "derece": {
                    "type": "string"
                },
                "dogumYeri": {
                    "type": "string"
                },
                "ekGosterge": {
                    "type": "string"
                },
                "eposta": {
                    "type": "string"
                },
                "girisTarihi": {
                    "type": "string"
                },
                "gorevAnaBilimDali": {
                    "type": "string"
                },
                "gorevBirim": {
                    "type": "string"
                },
                "gorevBolum": {
                    "type": "string"
                },
                "kadroAnaBilimDali": {
                    "type": "string"
                },
                "kadroBirim": {
                    "type": "string"
                },
                "kadroBolum": {
                    "type": "string"
                },
                "kadroBosalmaTarihi": {
                    "type": "string"
                },
                "kadrounvani": {
                    "type": "string"
                },
                "sicilNo": {
                    "type": "string"
                },
                "soyadi": {
                    "type": "string"
                },
                "tcKimlikNo": {
                    "type": "string"
                },
                "telefon": {
                    "type": "string"
                },
                "unvan": {
                    "type": "string"
                }
            }
        },
        "Models.PersonelSrvcResult": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Models.Personel"
                    }
                },
                "message": {
                    "type": "string"
                },
                "result": {
                    "type": "boolean"
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
	Version:     "1.0",
	Host:        "localhost:8011",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Swagger Bap Integration API",
	Description: "Integration için Data Sağlayan servis",
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