package gou

import "os"

// GetEnv reads the env variable specified by "key", or set a default value if the var not found
//
// Parameters:
//   - key: string - The env variable name, to get.
//   - defaultValue: string - The default value to return, in case key doens't exist.
//
// Returns:
//   - string - The string value or it's default value
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

// StrSlice slices a given string "s", from 1st char until the given "n"'th char
//
// Parameters:
//   - s: string - The source string.
//   - n: int - The number of chars to include.
//
// Returns:
//   - string - The resulting sub-string.
func StrSlice(s string, n int) string {
	strAsRune := []rune(s)

	if n < 0 || n >= len(strAsRune) {
		return s
	}

	return string(strAsRune[0:n])
}

// SliceContainsString checks if a given string exists in a strings slice.
//
// Parameters:
//   - slice: []string - The source slice where to check for the existence of a specific string.
//   - str: string - The string to check.
//
// Returns:
//   - bool - True if the string is found in the slice.
func SliceContainsString(slice []string, str string) bool {
	for _, el := range slice {
		if el == str {
			return true
		}
	}
	return false
}
