package utils

import (
	"fmt"
	"os"
	"regexp"
)

// For Config Load, replace Environment Variables found
func ReplaceEnvInConfig(vlr string) string {
	body := []byte(vlr)
	r := regexp.MustCompile(`\$\{([^{}]+)\}`)
	replaced := r.ReplaceAllFunc(body, func(b []byte) []byte {
		group1 := r.ReplaceAllString(string(b), `$1`)

		envValue := os.Getenv(group1)
		if len(envValue) > 0 {
			return []byte(envValue)
		}
		panic(fmt.Sprintf("Environment variable $%s was not set!!", group1))
	})
	return string(replaced)
}
