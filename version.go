package main

import (
	"fmt"
	"strings"

	release "github.com/kingdonb/footloose/pkg/version"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print footloose version",
	Run:   showVersion,
}

func init() {
	footloose.AddCommand(versionCmd)
}

var version = "git"

func showVersion(cmd *cobra.Command, args []string) {
	fmt.Println("version:", version)
	if version == "git" {
		return
	}
	release, err := release.FindLastRelease()
	if err != nil {
		fmt.Println("version: failed to check for new versions. You may want to check yourself at https://github.com/kingdonb/footloose/releases.")
		return
	}
	if strings.Compare(version, *release.TagName) != 0 {
		fmt.Printf("New version %v is available. More information at: %v\n", *release.TagName, *release.HTMLURL)
	}
}
