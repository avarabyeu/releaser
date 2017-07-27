package git

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

//ExecCmd executes command line command
func ExecCmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Executing command: %s %s\n", name, strings.Join(args, " "))
		return "", err
	}
	fmt.Printf("Executing command: %s %s\n", name, strings.Join(args, " "))
	res := string(out)
	if "" != res {
		fmt.Printf("%s\n", out)
	}
	return string(out), nil
}
