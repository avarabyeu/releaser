package util

import (
	"log"
	"testing"
)

func TestSubstrAfterLast(t *testing.T) {
	name := SubstrAfterLast("./release/releaser_linux_ppc64le", "/")
	filename := "releaser_linux_ppc64le"
	if filename != name {
		t.Errorf("Incorrect trim result. Expected %s, Actual: %s", filename, name)
	}
}

func TestPathInHome(t *testing.T) {
	//TODO
	log.Println(PathInHome(".git/"))
}
