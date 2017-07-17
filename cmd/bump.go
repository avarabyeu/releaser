package cmd

import (
	"fmt"
	"github.com/avarabyeu/releaser/util"
	"github.com/spf13/cobra"
	"log"
)

var bumpCommand = &cobra.Command{
	Use:   "bump",
	Short: "Bump new version number",
	Long:  `Bump new version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(bump(cmd))
	},
}

//GetSemverFile obtains path to semver file or returns default one
func GetSemverFile(cmd *cobra.Command) string {
	return cmd.Flag("file").Value.String()
}

//GetSemver returns Semver struct based on path provided
func GetSemver(cmd *cobra.Command) *util.Semver {
	ver := util.Load(GetSemverFile(cmd))
	return ver
}

//Bump bumps new provided version or next one. Saves result in version file
func bump(cmd *cobra.Command) string {
	newVersion := cmd.Flag("version").Value.String()
	ver := GetSemver(cmd)
	bump := ver.Bump(newVersion)
	err := ver.Save()
	if nil != err {
		log.Fatal("Cannot bump new version", err.Error())
	}
	return bump
}
