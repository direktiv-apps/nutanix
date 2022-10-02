// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Nutanix API access",
    "title": "nutanix",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "infrastructure"
      ],
      "container": "gcr.io/direktiv/functions/nutanix",
      "issues": "https://github.com/direktiv-apps/nutanix/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function gives access to Nutanix API v3. PLease use the Nutanix's API Explorer to see the available API commands. ",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/nutanix"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "auth",
                "api"
              ],
              "properties": {
                "api": {
                  "type": "object",
                  "required": [
                    "path"
                  ],
                  "properties": {
                    "body": {
                      "additionalProperties": false
                    },
                    "method": {
                      "description": "HTTP method to use",
                      "type": "string",
                      "default": "GET",
                      "example": "POST"
                    },
                    "path": {
                      "description": "API path to access",
                      "type": "string",
                      "example": "/vms/list"
                    }
                  }
                },
                "auth": {
                  "type": "object",
                  "required": [
                    "host",
                    "username",
                    "password"
                  ],
                  "properties": {
                    "host": {
                      "type": "string",
                      "example": "https://myserver:9440"
                    },
                    "password": {
                      "description": "Nutanix password",
                      "type": "string"
                    },
                    "skipVerify": {
                      "description": "Skip SSL certificate verification",
                      "type": "boolean",
                      "default": false
                    },
                    "username": {
                      "description": "Nutanix username",
                      "type": "string"
                    }
                  }
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Nutanix API response.",
            "schema": {
              "type": "object",
              "properties": {
                "nutanix": {
                  "type": "object",
                  "additionalProperties": false
                }
              }
            },
            "examples": {
              "nutanix": {
                "api_version": "3.1",
                "entities": [
                  {
                    "metadata": {
                      "categories": {},
                      "categories_mapping": {},
                      "creation_time": "2021-12-02T12:37:54Z",
                      "entity_version": "22",
                      "kind": "vm",
                      "last_update_time": "2022-08-05T10:40:01Z",
                      "owner_reference": {
                        "kind": "user",
                        "name": "admin",
                        "uuid": "00000000-0000-0000-0000-000000000000"
                      }
                    }
                  }
                ]
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "http",
              "data": {
                "kind": "string",
                "value": "{{ .API.Body | toJson }}"
              },
              "debug": true,
              "headers": [
                {
                  "Content-Type": "application/json"
                },
                {
                  "Accept": "application/json"
                }
              ],
              "insecure": "{{ if .Auth.SkipVerify }}true{{ else }}false{{ end }}",
              "method": "{{ if .API.Method }}{{ .API.Method }}{{ else }}GET{{ end }}",
              "password": "{{ .Auth.Password }}",
              "url": "{{ .Auth.Host }}/api/nutanix/v3{{ .API.Path }}",
              "username": "{{ .Auth.Username }}"
            }
          ],
          "output": "{\n  \"nutanix\": {{ (index (index . 0) \"result\") | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: nutanix\n  type: action\n  action:\n    function: nutanix\n    secrets: [\"nutanixpwd\"]\n    input: \n      auth:\n        host: https://myserver:9440\n        password: jq(.secrets.nutanixpwd)\n        username: myuser@nutanix\n        skipVerify: true\n      api:\n        path: \"/vms/list\"\n        method: POST\n        body:\n          kind: vm\n  catch:\n  - error: \"*\"",
            "title": "Basic"
          }
        ],
        "x-direktiv-function": "functions:\n- id: nutanix\n  image: gcr.io/direktiv/functions/nutanix:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "Password for the Nutanix user",
            "name": "nutanixpwd"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Nutanix API access",
    "title": "nutanix",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "infrastructure"
      ],
      "container": "gcr.io/direktiv/functions/nutanix",
      "issues": "https://github.com/direktiv-apps/nutanix/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function gives access to Nutanix API v3. PLease use the Nutanix's API Explorer to see the available API commands. ",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/nutanix"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/postParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Nutanix API response.",
            "schema": {
              "$ref": "#/definitions/postOKBody"
            },
            "examples": {
              "nutanix": {
                "api_version": "3.1",
                "entities": [
                  {
                    "metadata": {
                      "categories": {},
                      "categories_mapping": {},
                      "creation_time": "2021-12-02T12:37:54Z",
                      "entity_version": "22",
                      "kind": "vm",
                      "last_update_time": "2022-08-05T10:40:01Z",
                      "owner_reference": {
                        "kind": "user",
                        "name": "admin",
                        "uuid": "00000000-0000-0000-0000-000000000000"
                      }
                    }
                  }
                ]
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "http",
              "data": {
                "kind": "string",
                "value": "{{ .API.Body | toJson }}"
              },
              "debug": true,
              "headers": [
                {
                  "Content-Type": "application/json"
                },
                {
                  "Accept": "application/json"
                }
              ],
              "insecure": "{{ if .Auth.SkipVerify }}true{{ else }}false{{ end }}",
              "method": "{{ if .API.Method }}{{ .API.Method }}{{ else }}GET{{ end }}",
              "password": "{{ .Auth.Password }}",
              "url": "{{ .Auth.Host }}/api/nutanix/v3{{ .API.Path }}",
              "username": "{{ .Auth.Username }}"
            }
          ],
          "output": "{\n  \"nutanix\": {{ (index (index . 0) \"result\") | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: nutanix\n  type: action\n  action:\n    function: nutanix\n    secrets: [\"nutanixpwd\"]\n    input: \n      auth:\n        host: https://myserver:9440\n        password: jq(.secrets.nutanixpwd)\n        username: myuser@nutanix\n        skipVerify: true\n      api:\n        path: \"/vms/list\"\n        method: POST\n        body:\n          kind: vm\n  catch:\n  - error: \"*\"",
            "title": "Basic"
          }
        ],
        "x-direktiv-function": "functions:\n- id: nutanix\n  image: gcr.io/direktiv/functions/nutanix:1.0\n  type: knative-workflow",
        "x-direktiv-secrets": [
          {
            "description": "Password for the Nutanix user",
            "name": "nutanixpwd"
          }
        ]
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    },
    "postOKBody": {
      "type": "object",
      "properties": {
        "nutanix": {
          "type": "object",
          "additionalProperties": false
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBody": {
      "type": "object",
      "required": [
        "auth",
        "api"
      ],
      "properties": {
        "api": {
          "$ref": "#/definitions/postParamsBodyApi"
        },
        "auth": {
          "$ref": "#/definitions/postParamsBodyAuth"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBodyApi": {
      "type": "object",
      "required": [
        "path"
      ],
      "properties": {
        "body": {
          "additionalProperties": false
        },
        "method": {
          "description": "HTTP method to use",
          "type": "string",
          "default": "GET",
          "example": "POST"
        },
        "path": {
          "description": "API path to access",
          "type": "string",
          "example": "/vms/list"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBodyAuth": {
      "type": "object",
      "required": [
        "host",
        "username",
        "password"
      ],
      "properties": {
        "host": {
          "type": "string",
          "example": "https://myserver:9440"
        },
        "password": {
          "description": "Nutanix password",
          "type": "string"
        },
        "skipVerify": {
          "description": "Skip SSL certificate verification",
          "type": "boolean",
          "default": false
        },
        "username": {
          "description": "Nutanix username",
          "type": "string"
        }
      },
      "x-go-gen-location": "operations"
    }
  }
}`))
}
