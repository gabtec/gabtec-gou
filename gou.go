package gou

import "os"

// GetEnv - reads the env variable specified by "key", or set a default value if the var not found
func GetEnv(key, alternative string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return alternative
}