package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadBeefData(t *testing.T) {
	// Create a temporary file for testing
	testContent, err := os.ReadFile("./data.txt")
	if err != nil {
		t.Fatalf("Failed to read test data file: %v", err)
	}
	tmpFile := "test_data.txt"
	os.WriteFile(tmpFile, []byte(testContent), 0644)
	defer os.Remove(tmpFile) // Cleanup after test

	// Load data
	result := LoadBeefData(tmpFile)

	// Expected word counts
	expected := map[string]int{
		"eu":            44,
		"porchetta":     50,
		"quis":          40,
		"fatback":       38,
		"minim":         40,
		"burgdoggen":    34,
		"aliqua":        35,
		"ribeye":        40,
		"ea":            37,
		"deserunt":      49,
		"anim":          36,
		"cupim":         41,
		"mignon":        44,
		"tip":           49,
		"in":            144,
		"landjaeger":    40,
		"tempor":        38,
		"officia":       34,
		"enim":          41,
		"mollit":        26,
		"doner":         48,
		"excepteur":     31,
		"hamburger":     42,
		"leberkas":      20,
		"frankfurter":   43,
		"esse":          31,
		"cupidatat":     31,
		"kielbasa":      41,
		"andouille":     45,
		"consequat":     36,
		"cillum":        32,
		"chuck":         42,
		"reprehenderit": 42,
		"meatloaf":      33,
		"sunt":          37,
		"voluptate":     50,
		"laboris":       30,
		"venison":       39,
		"flank":         44,
		"dolor":         35,
		"short":         73,
		"adipisicing":   27,
		"elit":          40,
		"occaecat":      49,
		"ground":        43,
		"strip":         39,
		"meatball":      24,
		"irure":         47,
		"qui":           36,
		"bacon":         47,
		"drumstick":     29,
		"alcatra":       48,
		"shank":         42,
		"round":         43,
		"ut":            131,
		"id":            56,
		"spare":         43,
		"salami":        44,
		"laborum":       42,
		"tenderloin":    32,
		"belly":         35,
		"pastrami":      43,
		"prosciutto":    39,
		"hock":          29,
		"labore":        45,
		"filet":         44,
		"do":            36,
		"et":            37,
		"shankle":       34,
		"culpa":         43,
		"corned":        35,
		"sint":          44,
		"chop":          39,
		"turkey":        45,
		"dolore":        79,
		"nulla":         39,
		"ribs":          116,
		"rump":          37,
		"beef":          114,
		"jowl":          42,
		"shoulder":      34,
		"aute":          41,
		"tail":          34,
		"veniam":        33,
		"t-bone":        39,
		"pig":           31,
		"duis":          39,
		"eiusmod":       41,
		"ham":           71,
		"pancetta":      37,
		"swine":         41,
		"tri-tip":       35,
		"brisket":       37,
		"loin":          77,
		"chislic":       39,
		"ad":            37,
		"capicola":      36,
		"pariatur":      38,
		"aliquip":       49,
		"nostrud":       44,
		"biltong":       35,
		"steak":         39,
		"ex":            34,
		"kevin":         36,
		"incididunt":    34,
		"sirloin":       39,
		"cow":           41,
		"fugiat":        39,
		"chicken":       40,
		"exercitation":  43,
		"non":           38,
		"sausage":       47,
		"buffalo":       51,
		"pork":          138,
		"bresaola":      45,
		"magna":         31,
		"est":           35,
		"commodo":       49,
		"picanha":       34,
		"boudin":        43,
		"ball":          49,
		"consectetur":   52,
		"tongue":        24,
		"lorem":         30,
		"turducken":     47,
		"sed":           44,
		"jerky":         42,
		"ullamco":       46,
		"proident":      40,
		"ipsum":         46,
		"velit":         49,
		"nisi":          46,
	}

	// Check if all expected keys are present and have correct values
	for key, expectedValue := range expected {
		value, exists := result[key]
		assert.True(t, exists, "Expected key %s to exist", key)
		assert.Equal(t, expectedValue, value, "Expected value for key %s to be %d, got %d", key, expectedValue, value)
	}

	// Check if there are no unexpected keys
	for key := range result {
		_, exists := expected[key]
		assert.True(t, exists, "Unexpected key %s found", key)
	}
}
