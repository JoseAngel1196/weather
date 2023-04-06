/*
Copyright © 2023 Jose Hidalgo
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "weather",
	Short: "Weather report at your fingertips",
	Long:  `The weather CLI is a command-line tool that allows users to get current weather data for any location.`,
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
	cobra.OnInitialize(initConfig)

	persistent := rootCmd.PersistentFlags()
	persistent.String("location", "", "Your current location")
}

// initConfig reads a config file if set.
func initConfig() {
	dir, err := os.Getwd()
	cobra.CheckErr(err)

	viper.AddConfigPath(dir)
	viper.SetConfigType("toml")
	viper.SetConfigName("weather")
	viper.SafeWriteConfig()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

}
