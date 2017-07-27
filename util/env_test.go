package util

import (
	"testing"
	"reflect"
)

func TestGetEnvVars(t *testing.T) {
	vars := GetEnvVars()

	if "" == vars["PATH"] {
		t.Error("Path variable not found")
	}
}

func TestMapValues(t *testing.T) {
	mp := map[string]string{"key1": "value1", "key2": "value2"}
	values := MapValues(mp)

	if !reflect.DeepEqual(values, []string{"value1", "value2"}) {
		t.Error("Incorrect map processing")
	}
}
