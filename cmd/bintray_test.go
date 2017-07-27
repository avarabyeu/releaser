package cmd

import (
	"os"
	"testing"
	"text/template"
)

func TestUrlFormatter(t *testing.T) {

	fmt := "https://api.bintray.com/packages/{{.subject}}/{{.repo}}/{{.package}}/{{.versions}}"
	temp, _ := template.New("xxx").Parse(fmt)
	data := map[string]string{"subject": "subj"}
	temp.Execute(os.Stdout, data)

}
