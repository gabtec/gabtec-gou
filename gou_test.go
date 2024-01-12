package gou

import (
	"encoding/json"
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name           string
		envKey         string
		envVal         string
		envAlt         string
		expectedResult string
		prep           func()
		clean          func()
		// hasError bool
	}{
		// Test...
		{
			name:           "Should return the env var value",
			envKey:         "keyKey",
			envVal:         "keyValue",
			envAlt:         "altValue",
			expectedResult: "keyValue",
			prep:           func() { os.Setenv("keyKey", "keyValue") },
			clean:          func() { os.Unsetenv("keyKey") },
		},
		// Test...
		{
			name:           "Should return the env alternative value",
			envKey:         "keyKey",
			envVal:         "keyValue",
			envAlt:         "altValue",
			expectedResult: "altValue",
			prep:           func() {},
			clean:          func() {},
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

func TestStrSlice(t *testing.T) {
	tests := []struct {
		name           string // test name
		givenString    string // the string to process
		sliceSize      int
		expectedResult string
		prep           func()
		clean          func()
		// hasError bool
	}{
		// Test...
		{
			name:           "Should receive 'abcd' and return 'ab' if n=2 ",
			givenString:    "abcd",
			sliceSize:      2,
			expectedResult: "ab",
			prep:           func() { return },
			clean:          func() { return },
		},
		// Test...
		{
			name:           "Should receive 'abcd' and return 'abcd' if n bigger then str size ",
			givenString:    "abcd",
			sliceSize:      10,
			expectedResult: "abcd",
			prep:           func() { return },
			clean:          func() { return },
		},
		// Test...
		{
			name:           "Should receive 'abcd' and return 'abcd' if n is negative number ",
			givenString:    "abcd",
			sliceSize:      -10,
			expectedResult: "abcd",
			prep:           func() { return },
			clean:          func() { return },
		},
		// Test...
		{
			name:           "Should receive 'abcd' and return empty string if n=0 ",
			givenString:    "abcd",
			sliceSize:      0,
			expectedResult: "",
			prep:           func() { return },
			clean:          func() { return },
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prep()

			result := StrSlice(tt.givenString, tt.sliceSize)

			tt.clean()

			if result != tt.expectedResult {
				t.Errorf("Expected %s, got %s.", tt.expectedResult, result)
			}
		})
	}
}

// ---
func TestSliceContainsString(t *testing.T) {
	tests := []struct {
		name        string // test name
		srcSlice    []string
		checkStr    string // the string to process
		expectedRes bool
	}{
		// Test...
		{
			name: "Should return true when slice contains the string",
			srcSlice: []string{
				"audi",
				"vw",
				"mercedes",
			},
			checkStr:    "mercedes",
			expectedRes: true,
		},
		// Test...
		{
			name: "Should return false when slice does not contains the string",
			srcSlice: []string{
				"audi",
				"vw",
				"mercedes",
			},
			checkStr:    "notcontained",
			expectedRes: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SliceContainsString(tt.srcSlice, tt.checkStr)

			if result != tt.expectedRes {
				t.Errorf("Expected %v, got %v.", tt.expectedRes, result)
			}
		})
	}
}

// MergeMaps() tests
// ===========================================
func TestMergeMaps(t *testing.T) {
	tests := []struct {
		name        string // test name
		firstMap    map[string]interface{}
		secondMap   map[string]interface{}
		expectedRes map[string]interface{}
	}{
		// Test...
		{
			name: "Should work ok",
			firstMap: map[string]interface{}{
				"annot": "this is a annotation",
			},
			secondMap: map[string]interface{}{
				"todo": map[string]string{
					"monday": "buy milk",
				},
			},
			expectedRes: map[string]interface{}{
				"annot": "this is a annotation",
				"todo": map[string]string{
					"monday": "buy milk",
				},
			},
		},
		// Test...
		{
			name: "Should override existing keys",
			firstMap: map[string]interface{}{
				"annot": "this is a annotation",
				"desc":  "made for dev",
			},
			secondMap: map[string]interface{}{
				"desc": "made for testing",
				"todo": map[string]string{
					"monday": "buy milk",
				},
			},
			expectedRes: map[string]interface{}{
				"annot": "this is a annotation",
				"desc":  "made for testing",
				"todo": map[string]string{
					"monday": "buy milk",
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeMaps(tt.firstMap, tt.secondMap)

			// checkedKeys := []string{}

			// // secondMap first, because of override
			// for k := range tt.secondMap {
			// 	checkedKeys = append(checkedKeys, k)
			// 	t.Errorf("Expected %s, got %v.", tt.expectedRes[k], result[k])
			// }

			// // may not have all keys, because of Override
			// for k := range tt.firstMap {
			// 	if SliceContainsString(checkedKeys, k) {
			// 		continue
			// 	}
			// 	t.Errorf("Expected %s, got %v.", tt.expectedRes[k], result[k])
			// }

			jRes, _ := json.Marshal(result)
			jExp, _ := json.Marshal(tt.expectedRes)
			if string(jRes) != string(jExp) {
				t.Errorf("Expected %s, got %s.", jExp, jRes)
			}
		})
	}
}

// MergeStringMaps() tests
// ===========================================
func TestMergeStringMaps(t *testing.T) {
	tests := []struct {
		name        string // test name
		firstMap    map[string]string
		secondMap   map[string]string
		expectedRes map[string]string
	}{
		// Test...
		{
			name: "Should work ok",
			firstMap: map[string]string{
				"annot": "this is a annotation",
			},
			secondMap: map[string]string{
				"todo": "another anot",
			},
			expectedRes: map[string]string{
				"annot": "this is a annotation",
				"todo":  "another anot",
			},
		},
		// Test...
		{
			name: "Should override existing keys",
			firstMap: map[string]string{
				"annot": "this is a annotation",
				"desc":  "made for dev",
			},
			secondMap: map[string]string{
				"desc": "made for testing",
				"todo": "another anot",
			},
			expectedRes: map[string]string{
				"annot": "this is a annotation",
				"desc":  "made for testing",
				"todo":  "another anot",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeStringMaps(tt.firstMap, tt.secondMap)

			jRes, _ := json.Marshal(result)
			jExp, _ := json.Marshal(tt.expectedRes)
			if string(jRes) != string(jExp) {
				t.Errorf("Expected %s, got %s.", jExp, jRes)
			}
		})
	}
}
