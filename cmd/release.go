package cmd

import (
	"github.com/avarabyeu/releaser/git"
	"github.com/spf13/cobra"
	"log"
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

		//commit changes
		err := pushVersionFile(verFile, "[Releaser] Update release version")
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
		err = pushVersionFile(verFile, "[Releaser] Bump new snapshot version")
		if nil != err {
			log.Fatal("Cannot push new version")
		}

	},
}

func pushVersionFile(file string, message string) error {
	//commit changes
	err := git.Add(file)
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
}
