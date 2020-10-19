package commands

import (
	"fmt"

	"github.com/eankeen/bamboo/pkg/config"
	"github.com/safinsingh/stat"
	"github.com/spf13/cobra"
)

var location string

var rootCmd = &cobra.Command{
	Use:   "bamboo",
	Short: "A highly configurable bar for the X Window System",
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.Parse(location)
		fmt.Printf("%+v\n", conf)
	},
}

func init() {
	rootCmd.Flags().StringVarP(&location, "config", "c", "bamboo.toml", "")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		stat.Fail("Failed to execute root command")
	}
}
