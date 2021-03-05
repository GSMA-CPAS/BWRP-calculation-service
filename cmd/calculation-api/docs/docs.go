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
        "contact": {
            "name": "BWRP",
            "email": "developers@horizon.red"
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
        "/calculate": {
            "post": {
                "description": "Calculate the deal value by getting the contract and usage data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Calculate the dealvalue",
                "parameters": [
                    {
                        "description": "Discount agreements and usage data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CalculateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Result"
                        }
                    }
                }
            }
        },
        "/status": {
            "get": {
                "description": "Provides the version and hash of the executable code of the calculation engine",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Provide the status of the calculation engine",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Header"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CalculateRequest": {
            "type": "object",
            "properties": {
                "discounts": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/models.DiscountModel"
                    }
                },
                "usage": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.UsageData"
                    }
                }
            }
        },
        "models.Condition": {
            "type": "object",
            "properties": {
                "commitment": {
                    "$ref": "#/definitions/models.SelectedCondition"
                },
                "kind": {
                    "type": "string"
                }
            }
        },
        "models.DiscountModel": {
            "type": "object",
            "properties": {
                "condition": {
                    "$ref": "#/definitions/models.Condition"
                },
                "serviceGroups": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ServiceGroup"
                    }
                }
            }
        },
        "models.Header": {
            "type": "object",
            "properties": {
                "md5hash": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "models.IntermediateResult": {
            "type": "object",
            "properties": {
                "dealValue": {
                    "type": "string"
                },
                "homeTadigs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "service": {
                    "type": "string"
                },
                "visitorTadigs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.Pricing": {
            "type": "object",
            "properties": {
                "ratingPlan": {
                    "$ref": "#/definitions/models.RatingPlan"
                },
                "unit": {
                    "type": "integer"
                }
            }
        },
        "models.Rate": {
            "type": "object",
            "properties": {
                "fixedPrice": {
                    "type": "string"
                },
                "linearPrice": {
                    "type": "string"
                },
                "thresholds": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Tier"
                    }
                }
            }
        },
        "models.RatingPlan": {
            "type": "object",
            "properties": {
                "balancedRate": {
                    "$ref": "#/definitions/models.Rate"
                },
                "kind": {
                    "type": "string"
                },
                "rate": {
                    "$ref": "#/definitions/models.Rate"
                },
                "unbalancedRate": {
                    "$ref": "#/definitions/models.Rate"
                }
            }
        },
        "models.Result": {
            "type": "object",
            "properties": {
                "header": {
                    "$ref": "#/definitions/models.Header"
                },
                "intermediateResults": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.IntermediateResult"
                    }
                }
            }
        },
        "models.SelectedCondition": {
            "type": "object",
            "properties": {
                "currency": {
                    "type": "string"
                },
                "includingTaxes": {
                    "type": "boolean"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "models.Service": {
            "type": "object",
            "properties": {
                "accessPricing": {
                    "$ref": "#/definitions/models.Pricing"
                },
                "includedInCommitment": {
                    "type": "boolean"
                },
                "service": {
                    "type": "string"
                },
                "usagePricing": {
                    "$ref": "#/definitions/models.Pricing"
                }
            }
        },
        "models.ServiceGroup": {
            "type": "object",
            "properties": {
                "homeTadigs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "services": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Service"
                    }
                },
                "visitorTadigs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.Tier": {
            "type": "object",
            "properties": {
                "fixedPrice": {
                    "type": "string"
                },
                "linearPrice": {
                    "type": "string"
                },
                "start": {
                    "type": "string"
                }
            }
        },
        "models.UsageData": {
            "type": "object",
            "properties": {
                "charges": {
                    "type": "string"
                },
                "homeTadig": {
                    "type": "string"
                },
                "service": {
                    "type": "string"
                },
                "taxes": {
                    "type": "string"
                },
                "units": {
                    "type": "string"
                },
                "usage": {
                    "type": "string"
                },
                "visitorTadig": {
                    "type": "string"
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
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{"http", "https"},
	Title:       "Calculation API",
	Description: "Calculation API",
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
