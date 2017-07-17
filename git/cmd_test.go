package git

import (
	"log"
	"strings"
	"testing"
)

func TestExecCmd(t *testing.T) {
	_, err := ExecCmd("ls", "-la")

	if nil != err {
		t.Errorf("Execution failed: %s", err)
	}
}

func TestExecCmd2(t *testing.T) {
	out, err := ExecCmd("git", "tag", "-l")

	if nil != err {
		t.Errorf("Execution failed: %s", err)
	}

	parts := strings.Split(out, "\n")
	for i, part := range parts {
		log.Printf("%d: %s", i, part)
	}
}
