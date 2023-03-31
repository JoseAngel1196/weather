package config

import (
	"os"

	"github.com/JoseAngel1196/weather/print"
)

func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		print.ExitOnError("Hey, you need to set the credentials on your .env 😭", nil)
	}

	return value
}
