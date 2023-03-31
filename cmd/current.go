/*
Copyright © 2023 Jose Hidalgo
*/
package cmd

import (
	"github.com/JoseAngel1196/weather/pkg"
	"github.com/spf13/cobra"
)

// currentCmd represents the current command
var currentCmd = &cobra.Command{
	Use:     "current",
	Short:   "Show current weather data for a location, including temperature, humidity, wind speed, and more.",
	Long:    `...`,
	Run:     pkg.Current,
	Aliases: []string{"c", "current"},
}

func init() {
	rootCmd.AddCommand(currentCmd)
}
