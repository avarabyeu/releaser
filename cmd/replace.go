package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"log"
	"path"
	"io/ioutil"
	"text/template"
	"bytes"
	"github.com/avarabyeu/releaser/util"
)

var replaceCommand = &cobra.Command{
	Use:   "replace",
	Short: "Replaces placeholders in files",
	Long:  `Replaces placeholders in files`,
	Run: func(cmd *cobra.Command, args []string) {
		err := replace(cmd, config.Replace)
		if nil != err {
			log.Fatalf("Cannot execute replace command: %s", err)
		}
	},
}

//Replaces placeholders in files
func replace(cmd *cobra.Command, files []string) error {
	wd, err := os.Getwd()
	if nil != err {
		log.Printf("Cannot get work dir! %s", err.Error())
		return err
	}

	data := util.GetEnvVars()
	data["version"] = GetSemver(cmd).Current()

	for _, repl := range files {
		replFile := path.Join(wd, repl)
		content, err := ioutil.ReadFile(replFile)
		if nil != err {
			log.Printf("Cannot read file to replace: %s", err.Error())
			return err
		}
		tmpl, _ := template.New(repl).Parse(string(content))
		tmplRes := new(bytes.Buffer)

		err = tmpl.Execute(tmplRes, data)
		if nil != err {
			return err
		}

		err = ioutil.WriteFile(replFile, tmplRes.Bytes(), 0644)
		if nil != err {
			return err
		}

	}

	return nil
}
