package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"min=18"`
}

func main() {
	user := User{
		Name: "Alice",
		Age:  20,
	}
	t := reflect.TypeOf(user)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		validateTag := field.Tag.Get("validate")
		fmt.Printf("Field: %s, JSON Tag: %s, Validate Tag: %s\n", field.Name, jsonTag, validateTag)
	}
}
