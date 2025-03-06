package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// LoadBeefData reads file and counts words
func LoadBeefData(filename string) map[string]int {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	text := string(content)
	words := strings.Fields(text) // Split words by spaces

	// Count word occurrences
	wordCount := make(map[string]int)
	for _, word := range words {
		word = strings.ToLower(strings.Trim(word, ".,")) // Normalize words
		wordCount[word]++
	}

	for key, val := range wordCount {
		fmt.Printf("\"%s\":   %d,\n", key, val)
	}

	log.Println("File loaded successfully with beef counts!")
	return wordCount
}
