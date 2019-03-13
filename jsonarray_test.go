package mappy

import "testing"

func TestJsonArray(t *testing.T) {
	data := `[{},{},{},{}]`
	jsonArray, err := JsonArrayFromString(data)
	if err != nil {
		t.Error(err)
	}

	if len(jsonArray) != 4 {
		t.Error("Should have been 4 objects")
	}
}

