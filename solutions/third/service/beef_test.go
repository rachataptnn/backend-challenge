package service

import (
	"os"
	"testing"
)

func setup() {
	content := `Beef is great. Beef is tasty.`
	os.WriteFile("../src/beef_test.txt", []byte(content), 0644)
}

func teardown() {
	os.Remove("../src/beef_test.txt")
}

func TestReadBeefFile(t *testing.T) {
	setup()
	defer teardown()

	err := ReadBeefFile("../src/beef_test.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := map[string]int{
		"beef":  2,
		"is":    2,
		"great": 1,
		"tasty": 1,
	}

	for k, v := range expected {
		if BeefCounts[k] != v {
			t.Errorf("Expected %d for word %s, got %d", v, k, BeefCounts[k])
		}
	}
}

func TestGetBeefSummary(t *testing.T) {
	setup()
	defer teardown()

	ReadBeefFile("../src/beef_test.txt")

	expected := map[string]int{
		"beef":  2,
		"is":    2,
		"great": 1,
		"tasty": 1,
	}

	summary := GetBeefSummary()
	for k, v := range expected {
		if summary[k] != v {
			t.Errorf("Expected %d for word %s, got %d", v, k, summary[k])
		}
	}
}
