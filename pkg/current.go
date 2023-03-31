package pkg

import (
	"github.com/JoseAngel1196/weather/internal"
	"github.com/spf13/cobra"
)

func Current(cmd *cobra.Command, args []string) {
	internal.GetUserLocation()
}
