package gou

import "os"

// GetEnv - reads the env variable specified by "key", or set a default value if the var not found
func GetEnv(key, alternative string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return alternative
}

// StrSlice - slices a given string "s", from 1st char until the given "n"'th char
func StrSlice(s string, n int) string {
	strAsRune := []rune(s)

	if n < 0 || n >= len(strAsRune) {
		return s
	}

	return string(strAsRune[0:n])
}