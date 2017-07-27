package git

import (
	"strings"
)

//GetLocalTags returns tags from local repository
func GetLocalTags() ([]string, error) {
	out, err := ExecCmd("git", "tag", "-l")

	if nil != err {
		return nil, err
	}
	if "" == out {
		return []string{}, nil
	}
	return strings.Split(out, "\n"), nil
}

//DeleteTag deletes local tag
func DeleteTag(tag string) error {
	_, err := ExecCmd("git", "tag", "-d", tag)
	return err
}

//FetchRemoteTags fetches remote tags
func FetchRemoteTags() error {
	_, err := ExecCmd("git", "fetch", "--tags")
	return err
}

//CreateTag creates tag
func CreateTag(tag, message string) error {
	_, err := ExecCmd("git", "tag", tag, "-m", message)
	return err
}

//Push pushes changes to origin
func Push(origin string) error {
	_, err := ExecCmd("git", "push", origin)
	return err
}

//PushTags pushes tags to origin
func PushTags() error {
	_, err := ExecCmd("git", "push", "--tags")
	return err
}

//Add adds path to commit list
func Add(path string) error {
	_, err := ExecCmd("git", "add", path)
	return err
}

//Commit commits with provided message
func Commit(message string) error {
	_, err := ExecCmd("git", "commit", "-m", message)
	return err
}
