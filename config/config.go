package config

import (
	"errors"
	"os"
)

func GetEnv(key string) (string, error) {
	value, ok := os.LookupEnv(key)

	if !ok {
		return "", errors.New((".env is not setup! 😭"))
	}

	return value, nil
}
