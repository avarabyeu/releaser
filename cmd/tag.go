package cmd

import (
	"fmt"
	"github.com/avarabyeu/releaser/git"
	"github.com/spf13/cobra"
	"log"
)

var tagCommand = &cobra.Command{
	Use:   "tag",
	Short: "Tags new version in git",
	Long:  `Print the current version number`,
	Run: func(cmd *cobra.Command, args []string) {

		tag := GetSemver(cmd).Current()
		newTag(tag)

	},
}

func newTag(tag string) {

	tags, err := git.GetLocalTags()

	if nil != err {
		log.Fatalf("Execution failed: %s", err)
	}

	for _, tag := range tags {
		git.DeleteTag(tag)
	}
	err = git.FetchRemoteTags()
	if nil != err {
		log.Fatalf("Cannot fetch remote tags: %s", err)
	}

	err = git.CreateTag(tag, fmt.Sprintf("Tag %s", tag))
	if nil != err {
		log.Fatalf("Cannot create tag: %s", err)
	}

	err = git.PushTags()
	if nil != err {
		log.Fatalf("Cannot push tag: %s", err)
	}

}
