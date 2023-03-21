/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"multi-clash-subscriber/internal/conf"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "generate config",
	Long:  `generate config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		output, err := cmd.Flags().GetString("output")
		if err != nil {
			return errors.Wrap(err, "get output flag failed")
		}
		if output == "" {
			return errors.New("output flag is required")
		}

		if err := conf.Generate(output); err != nil {
			return errors.Wrap(err, "generate config failed")
		}

		return nil
	},
}

func init() {
	genCmd.AddCommand(configCmd)
}
