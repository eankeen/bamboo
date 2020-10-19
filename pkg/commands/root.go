package commands

import (
	"fmt"

	"github.com/eankeen/bamboo/pkg/config"
	"github.com/eankeen/bamboo/pkg/x"
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
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.Parse(location)
		stat.Success("Deserialized configuration")
		if verbose {
			stat.Info(fmt.Sprintf("Parsed configuration: %+v", conf))
		}
		x.Draw(conf)
	},
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
