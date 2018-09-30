// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-09-29 16:42:22.855923839 +0800 CST m=+0.048320935

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/api/v1/admin/user/login": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Admin Login"
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Admin Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ]
            }
        },
        "/api/v1/user/login": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "User Register"
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ]
            }
        },
        "/api/v1/user/profile": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "User Profile"
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "uid",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    }
                ]
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
