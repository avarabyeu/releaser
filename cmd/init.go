package cmd

import (
	"fmt"
	"github.com/avarabyeu/releaser/util"
	"github.com/coreos/go-semver/semver"
	"github.com/spf13/cobra"
	"log"
)

const initialVersionDefault = "0.0.1"

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Creates version file",
	Long:  `Creates version file`,
	Run: func(cmd *cobra.Command, args []string) {

		versionStr := cmd.Flag("version").Value.String()
		if "" == versionStr {
			versionStr = initialVersionDefault
		}
		v, _ := semver.NewVersion(versionStr)

		file := cmd.Flag("file").Value.String()
		ver := util.New(file, v)
		err := ver.Save()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Bumped initial version to %s\n", v.String())
	},
}
