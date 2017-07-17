package util

import (
	"log"
	"os/user"
	"path"
	"strings"
)

//SubstrAfterLast returns substring after last occurrence of delimiter
func SubstrAfterLast(str string, delimiter string) string {
	ind := strings.LastIndex(str, delimiter)
	if -1 == ind {
		return ""
	}
	return string(str[ind+1:])
}

//GetUserHome returns path to user home directory
func GetUserHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

//PathInHome appends path to user home directory
func PathInHome(file string) string {
	return path.Join(GetUserHome(), file)
}
