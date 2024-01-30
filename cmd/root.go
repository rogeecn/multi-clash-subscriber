/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"multi-clash-subscriber/config"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:          "multi-clash-subscriber",
	Short:        "combine multi clash subscription",
	Long:         "combine multi clash subscription",
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.SetConfigType("toml")
		if cfgFile == "" {
			viper.SetConfigName("config.toml")
		} else {
			viper.SetConfigFile(cfgFile)
		}

		wd, err := os.Getwd()
		if err != nil {
			return errors.Wrap(err, "get work dir failed")
		}
		log.Printf("add config path: %s", wd+"/conf/")
		viper.AddConfigPath(wd + "/conf/")

		if err := viper.ReadInConfig(); err != nil {
			return errors.Wrap(err, "read in config failed")
		}

		if err := viper.Unmarshal(&config.C); err != nil {
			return errors.Wrap(err, "unmarshal config failed")
		}

		return nil
	},
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
