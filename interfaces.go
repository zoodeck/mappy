package mappy

import (
	"encoding/json"
)

type JsonObject interface {
	Contains(key string) bool
	Int(key string) (int, error)
	Bool(key string) (bool, error)
	String(key string) (string, error)
	JsonArray(key string) ([]JsonObject, error)
	JsonObject(key string) (JsonObject, error)
	Value() (interface{}, error)
	IsValue() bool
	JsonString() (string, error)
	Keys() ([]string, error)
}

type JsonArray []JsonObject

func JsonObjectFromBytes(data []byte) (JsonObject, error) {
	var innerMap map[string]interface{}
	mapError := json.Unmarshal(data, &innerMap)
	if mapError == nil {
		return innerJsonObject{innerMap, nil}, nil
	}

	var innerValue interface{}
	valueError := json.Unmarshal(data, &innerValue)
	if valueError == nil {
		return innerJsonObject{map[string]interface{}{}, innerValue}, nil
	}

	return nil, mapError
}

func JsonObjectFromString(data string) (JsonObject, error) {
	return JsonObjectFromBytes([]byte(data))
}

func JsonArrayFromBytes(data []byte) ([]JsonObject, error) {
	var innerArray []interface{}
	err := json.Unmarshal(data, &innerArray)
	if err != nil {
		return nil, err
	}

	jsonObjects := make([]JsonObject, len(innerArray))
	for i := 0; i < len(innerArray); i++ {
		bits, err := json.Marshal(innerArray[i])
		if err != nil {
			return nil, err
		}

		jsonObject, _ := JsonObjectFromBytes(bits)
		jsonObjects[i] = jsonObject
	}

	return jsonObjects, nil
}

func JsonArrayFromString(data string) ([]JsonObject, error) {
	return JsonArrayFromBytes([]byte(data))
}
