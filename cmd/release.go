package cmd

import (
	"github.com/avarabyeu/releaser/git"
	"github.com/spf13/cobra"
	"log"
	"strings"
	"github.com/avarabyeu/releaser/util"
)

var releaseCommand = &cobra.Command{
	Use:   "release",
	Short: "Release new version",
	Long:  `Release new version`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("v1")
		verFile := GetSemverFile(cmd)

		//remove snapshot
		bump(cmd)

		//replace if needed
		replace(cmd, config.Replace)
		toCommit := []string{}
		_ = append(toCommit, util.MapValues(config.Replace)...)
		_ = append(toCommit, verFile)

		//commit changes
		err := pushFiles("[Releaser] Update release version", toCommit...)
		if nil != err {
			log.Fatal("Cannot push new version")
		}

		//upload to bintray
		err = uploadToBintray(cmd)
		if nil != err {
			log.Fatal("Cannot upload artifacts to bintray")
		}

		//create tag
		semver := GetSemver(cmd)
		newTag(semver.Current())

		//bump new snapshot
		semver.NextSnapshot("")
		semver.Save()
		if nil != err {
			log.Fatal("Cannot upload artifacts to bintray")
		}

		//push changes
		err = pushFiles("[Releaser] Bump new snapshot version", verFile)
		if nil != err {
			log.Fatal("Cannot push new version")
		}

	},
}

func pushFiles(message string, files ...string) error {
	//commit changes
	err := git.Add(strings.Join(files, " "))
	if nil != err {
		return err
	}
	git.Commit(message)
	if nil != err {
		return err
	}

	git.Push("origin")
	if nil != err {
		return err
	}

	return nil
}
