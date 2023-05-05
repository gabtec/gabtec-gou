package gou

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name string
		envKey string 
		envVal string 
		envAlt string 
		expectedResult string
		prep func()
		clean func()
		// hasError bool
	}{
		// Test...
		{
			name: "Should return the env var value",
			envKey: "keyKey",
			envVal: "keyValue",
			envAlt: "altValue",
			expectedResult: "keyValue",
			prep: func() { os.Setenv("keyKey", "keyValue")},
			clean: func() { os.Unsetenv("keyKey")},
		},
		// Test...
		{
			name: "Should return the env alternative value",
			envKey: "keyKey",
			envVal: "keyValue",
			envAlt: "altValue",
			expectedResult: "altValue",
			prep: func() {},
			clean: func() {},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prep()

			result := GetEnv(tt.envKey, tt.envAlt)
		
			tt.clean()

			if result != tt.expectedResult {
				t.Errorf("Expected %s, got %s.", tt.expectedResult, result)
			}
		})
	}
}