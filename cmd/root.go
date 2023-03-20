package cmd

import (
	"log"
	"multi-clash-subscriber/config"
	"multi-clash-subscriber/internal/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "multi-clash-subscriber",
	Short: "A",
	Long:  `A.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigType("toml")
		wd, _ := os.Getwd()
		viper.AddConfigPath(wd)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal("1.", err)
		}

		var c config.Config
		if err := viper.Unmarshal(&c); err != nil {
			log.Fatal("2.", err)
		}

		if err := http.Serve(&c); err != nil {
			log.Fatal("3.", err)
		}
	},
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.toml", "config file (default is $HOME/.multi-clash-subscriber.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
