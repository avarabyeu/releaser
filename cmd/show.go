package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var showCommand = &cobra.Command{
	Use:   "show",
	Short: "Print the current version number",
	Long:  `Print the current version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(GetSemver(cmd).Version.String())
	},
}
