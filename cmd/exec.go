package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var exec = &cobra.Command{
	Use:   "exec",
	Short: "Executes bunch of commands",
	Long:  `Executes bunch of commands`,
	Run: func(cmd *cobra.Command, args []string) {
		commands := make(map[string]struct{}, len(RootCommand.Commands()))
		for _, c := range RootCommand.Commands() {
			commands[c.Name()] = struct{}{}
		}

		for _, arg := range args {
			if _, ok := commands[arg]; ok {
				RootCommand.SetArgs([]string{arg})
				_ = RootCommand.Execute()

			} else {
				log.Fatalf("No such command: %s", arg)

			}
		}
	},
}
