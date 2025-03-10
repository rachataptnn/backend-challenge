package service

import (
	"os"
	"strings"
)

var BeefCounts map[string]int

func ReadBeefFile(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	words := strings.Fields(string(content))

	m := make(map[string]int)
	for _, word := range words {
		clean := strings.Trim(word, ".,")
		word = strings.ToLower(clean)
		m[word]++
	}

	BeefCounts = m
	return nil
}

func GetBeefSummary() map[string]int {
	return BeefCounts
}
