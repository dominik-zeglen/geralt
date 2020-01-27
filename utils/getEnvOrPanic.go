package utils

import (
	"fmt"
	"os"
)

func GetEnvOrPanic(key string) string {
	r := os.Getenv(key)
	if r == "" {
		panic(fmt.Errorf("Environment variable %s not set", key))
	}

	return r
}
