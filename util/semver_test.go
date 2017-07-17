package util

import (
	"github.com/coreos/go-semver/semver"
	"testing"
)

func TestSemver_BumpDefault(t *testing.T) {
	v, _ := semver.NewVersion("0.0.1")
	s := Semver{file: "", Version: v}
	nv := s.Bump("")
	if "0.0.2" != nv {
		t.Error("Bump command fails")
	}

}

func TestBumpWithVersion(t *testing.T) {

	v, _ := semver.NewVersion("0.0.1")
	s := Semver{file: "", Version: v}
	nv := s.Bump("0.0.2")

	if "0.0.2" != nv {
		t.Error("Bump command fails")
	}

}
