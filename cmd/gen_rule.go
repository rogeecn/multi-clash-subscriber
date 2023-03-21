package cmd

import (
	"multi-clash-subscriber/internal/rule"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// ruleCmd represents the rule command
var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "generate rules from https://github.com/Loyalsoldier/clash-rules",
	Long:  "generate rules from https://github.com/Loyalsoldier/clash-rules",
	RunE: func(cmd *cobra.Command, args []string) error {
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return errors.Wrap(err, "get output flag failed")
		}

		if output == "" {
			return errors.New("output flag is required")
		}

		if err := rule.Generate(output); err != nil {
			return errors.Wrap(err, "generate rule failed")
		}

		return nil
	},
}

func init() {
	genCmd.AddCommand(ruleCmd)
}
