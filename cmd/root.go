package cmd

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

func init() {
	RootCommand.PersistentFlags().StringP("file", "f", "VERSION", "Version file to store current version")
	RootCommand.PersistentFlags().StringP("version", "v", "", "Version to be used")
	RootCommand.PersistentFlags().StringP("bintray.user", "", "", "BintrayConf user name")
	RootCommand.PersistentFlags().StringP("bintray.org", "", "", "BintrayConf organization")
	RootCommand.PersistentFlags().StringP("bintray.token", "", "", "BintrayConf token")
	RootCommand.PersistentFlags().StringP("bintray.repo", "", "", "BintrayConf repository")
	RootCommand.PersistentFlags().StringP("bintray.pack", "", "", "BintrayConf package")
	RootCommand.PersistentFlags().StringP("artifactsFolder", "", "release", "Folder with artifacts to upload")
	RootCommand.PersistentFlags().StringP("replace", "", "release", "Replaces placeholders in files")

	cobra.OnInitialize(initConfig)

	RootCommand.AddCommand(showCommand)
	RootCommand.AddCommand(initCommand)
	RootCommand.AddCommand(bumpCommand)
	RootCommand.AddCommand(tagCommand)
	RootCommand.AddCommand(bintrayCommand)
	RootCommand.AddCommand(exec)
	RootCommand.AddCommand(releaseCommand)
	RootCommand.AddCommand(replaceCommand)
}

var config *Config

//RootCommand entry point for all commands
var RootCommand = &cobra.Command{
	Use:   "releaser",
	Short: "Release assistant",
	Long:  `Release assistant built by avarabyeu`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func initConfig() {

	conf := viper.New()

	conf.BindPFlag("file", RootCommand.PersistentFlags().Lookup("file"))
	conf.BindPFlag("version", RootCommand.PersistentFlags().Lookup("version"))
	conf.BindPFlag("bintray.user", RootCommand.PersistentFlags().Lookup("bintray.user"))
	conf.BindPFlag("bintray.org", RootCommand.PersistentFlags().Lookup("bintray.org"))
	conf.BindPFlag("bintray.token", RootCommand.PersistentFlags().Lookup("bintray.token"))
	conf.BindPFlag("bintray.repo", RootCommand.PersistentFlags().Lookup("bintray.repo"))
	conf.BindPFlag("bintray.pack", RootCommand.PersistentFlags().Lookup("bintray.pack"))
	conf.BindPFlag("artifactsFolder", RootCommand.PersistentFlags().Lookup("artifactsFolder"))
	conf.BindPFlag("replace", RootCommand.PersistentFlags().Lookup("replace"))

	// Search config in home directory with name ".cobra" (without extension).
	conf.SetConfigName(".releaser")
	conf.AddConfigPath(".")
	conf.SetConfigType("yml")

	if err := conf.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	if err := conf.Unmarshal(&config); err != nil {
		fmt.Println("Can't unmarshall:", err)
		os.Exit(1)
	}

}

//Config represents project config
type Config struct {
	Bintray         *BintrayConf `mapstructure:"bintray"`
	ArtifactsFolder string
	Replace         map[string]string
}

//BintrayConf represents project config
type BintrayConf struct {
	User  string
	Token string
	Repo  string
	Pack  string
	Org   string
}

func (b *BintrayConf) getAuth() string {
	return "Basic " + b64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", b.User, b.Token)))
}
