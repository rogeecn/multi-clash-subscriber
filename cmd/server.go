/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"multi-clash-subscriber/internal/http"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start http server",
	Long:  `start http server`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return http.Serve()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
