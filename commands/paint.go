package commands

import (
	"encoding/json"

	"github.com/eankeen/bamboo/config"
	"github.com/eankeen/bamboo/x"
	"github.com/safinsingh/stat"
	"github.com/spf13/cobra"
)

var paintCmd = &cobra.Command{
	Use:   "paint",
	Short: "Paint the bar",
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.Parse(location, verbose)
		if err != nil {
			stat.Fail("Failed to unmarshal TOML data: " + err.Error())
		}
		stat.Success("Deserialized configuration")
		if verbose {
			json, _ := json.MarshalIndent(conf, "", "   ")
			stat.Info("Parsed configuration:\n", string(json))
		}
		x.Draw(conf)
	},
}

func init() {
	rootCmd.AddCommand(paintCmd)
}
