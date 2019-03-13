package mappy

import (
	"fmt"
	"testing"
)

func TestJsonObjectStringValue(t *testing.T) {
	data := `{"hello":"world"}`
	jsonObject, err := JsonObjectFromString(data)
	if err != nil {
		t.Error(err)
	}

	if !jsonObject.Contains("hello") {
		t.Error("Should contain `hello`")
	}

	value, err := jsonObject.String("hello")
	if err != nil {
		t.Error(err)
	}

	if value != "world" {
		t.Error("Value should equal `world`")
	}
}


func TestJsonObjectInvalidString(t *testing.T) {
	data := `{hello: "world"}`
	_, err := JsonObjectFromString(data)
	if err == nil {
		t.Error("Error should not have been nil")
	}
}

func TestJsonObjectComplicated(t *testing.T) {
	data := `
		{
			"string": "string",
			"int": 1,
			"bool": true,
			"object": {
				"hello": "world"
			},
			"array": [
				1,
				true,
				{"hello": "world"},
				"stringy"
			]
		}
	`

	jsonObject, err := JsonObjectFromString(data)
	if err != nil {
		t.Error(err)
	}

	jsonArray, err := jsonObject.JsonArray("array")
	if err != nil {
		t.Error(err)
	}

	for _, o := range jsonArray {
		if o.IsValue() {
			fmt.Println(o.Value())
		} else {
			fmt.Println(o.Keys())
		}
	}
}