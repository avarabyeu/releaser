package util

import (
	"testing"
)

func TestGetEnvVars(t *testing.T) {
	vars := GetEnvVars()

	if "" == vars["PATH"] {
		t.Error("Path variable not found")
	}
}
