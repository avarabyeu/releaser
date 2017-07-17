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
		pushVersionFile(verFile, "[Releaser] Update release version")

		//upload to bintray
		uploadToBintray(cmd)

		//create tag
		semver := GetSemver(cmd)
		newTag(semver.Current())

		//bump new snapshot
		semver.NextSnapshot("")
		semver.Save()

		//push changes
		pushVersionFile(verFile, "[Releaser] Bump new snapshot version")
	},
}

func pushVersionFile(file string, message string) {
	//commit changes
	git.Add(file)
	git.Commit(message)
	git.Push("origin")
}
