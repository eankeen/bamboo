package commands

import (
	"github.com/safinsingh/stat"
	"github.com/spf13/cobra"
)

var (
	location string
	verbose  bool
)

var rootCmd = &cobra.Command{
	Use:   "bamboo",
	Short: "A highly configurable bar for the X Window System",
}

func init() {
	rootCmd.Flags().StringVarP(&location, "config", "c", "bamboo.toml", "location of configuration file")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "run in verbose mode")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		stat.Fail("Failed to execute root command")
	}
}
