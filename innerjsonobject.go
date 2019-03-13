package mappy

import (
	"encoding/json"
	"fmt"
)

type innerJsonObject struct {
	innerMap map[string]interface{}
	innerValue interface{}
}

func (i innerJsonObject) Contains(key string) bool {
	_, contains := i.innerMap[key]
	return contains
}

func (i innerJsonObject) Int(key string) (int, error) {
	if !i.Contains(key) {
		return 0, DoesNotContainError
	}

	value := i.innerMap[key]

	switch value.(type) {
	case int:
		return value.(int), nil
	default:
		return 0, InvalidTypeError
	}
}

func (i innerJsonObject) Bool(key string) (bool, error) {
	if !i.Contains(key) {
		return false, DoesNotContainError
	}

	value := i.innerMap[key]

	switch value.(type) {
	case bool:
		return value.(bool), nil
	default:
		return false, InvalidTypeError
	}
}

func (i innerJsonObject) String(key string) (string, error) {
	if !i.Contains(key) {
		return "", DoesNotContainError
	}

	value := i.innerMap[key]

	switch value.(type) {
	case string:
		return value.(string), nil
	default:
		return "", InvalidTypeError
	}
}

func (i innerJsonObject) JsonArray(key string) ([]JsonObject, error) {
	if !i.Contains(key) {
		return nil, DoesNotContainError
	}

	value := i.innerMap[key]

	switch value.(type) {
	case []interface{}:
		fmt.Println("gets jsonarray here:", value)

		bits, err := json.Marshal(value)
		if err != nil {
			return nil, err
		}

		fmt.Println("value", string(bits))

		return JsonArrayFromBytes(bits)
	default:
		return nil, InvalidTypeError
	}
}

func (i innerJsonObject) JsonObject(key string) (JsonObject, error) {
	if !i.Contains(key) {
		return nil, DoesNotContainError
	}

	value := i.innerMap[key]

	switch value.(type) {
	case JsonObject:
		return value.(JsonObject), nil
	default:
		return nil, InvalidTypeError
	}
}

func (i innerJsonObject) Value() (interface{}, error) {
	if i.innerValue == nil {
		return nil, MissingValueError
	}

	return i.innerValue, nil
}

func (i innerJsonObject) IsValue() bool {
	return i.innerValue != nil
}

func (i innerJsonObject) JsonString() (string, error) {
	bits, err := json.Marshal(i.innerMap)
	if err != nil {
		return "", err
	}

	return string(bits), nil
}

func (i innerJsonObject) Keys() ([]string, error) {
	if i.innerValue != nil {
		return nil, ValuePresentError
	}

	keys := []string{}
	for key, _ := range i.innerMap {
		keys = append(keys, key)
	}

	return keys, nil
}