package pkg

import (
	"fmt"
	"os"

	"github.com/JoseAngel1196/weather/print"
)

func exitOnError(str string, err error) {
	fmt.Println("")
	if str != "" {
		print.Message(str, print.Error)
	}
	if err != nil {
		fmt.Printf("Error message: %s\n", err)
		print.Message("Unexpected error", print.Error)
	}
	os.Exit(1)
}
