package cmd

import (
	"bytes"
	"github.com/avarabyeu/releaser/util"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
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
func replace(cmd *cobra.Command, files map[string]string) error {
	wd, err := os.Getwd()
	if nil != err {
		log.Printf("Cannot get work dir! %s", err.Error())
		return err
	}

	data := util.GetEnvVars()
	data["version"] = GetSemver(cmd).Current()

	for tmplPath, resPath := range files {
		replFile := path.Join(wd, tmplPath)
		tmplContent, err := ioutil.ReadFile(replFile)
		if nil != err {
			log.Printf("Cannot read file to replace: %s", err.Error())
			return err
		}
		tmpl, _ := template.New(tmplPath).Parse(string(tmplContent))
		processed := new(bytes.Buffer)

		err = tmpl.Execute(processed, data)
		if nil != err {
			return err
		}

		err = ioutil.WriteFile(resPath, processed.Bytes(), 0644)
		if nil != err {
			return err
		}

	}

	return nil
}
