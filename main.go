/*
Copyright © 2023 Jose Hidalgo
*/
package main

import (
	"github.com/JoseAngel1196/weather/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Gather credentials from the environment
	godotenv.Load(".env")

	cmd.Execute()
}
