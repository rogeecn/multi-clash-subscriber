package cmd

import (
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate something",
	Long:  `generate config/rules`,
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.PersistentFlags().StringP("output", "o", "", "output file")
}
