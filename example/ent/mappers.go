// Code generated by entc, DO NOT EDIT.

package ent

import openapi "github.com/ogen-go/ent2ogen/example/openapi"

func (e *City) ToOpenAPI() openapi.City {
	return openapi.City{
		Name: e.Name,
	}
}

func (e *User) ToOpenAPI() openapi.User {
	return openapi.User{
		ID:        e.ID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Username:  e.Username,
	}
}
