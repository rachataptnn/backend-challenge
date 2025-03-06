package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"bufio"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetBeefSummary(t *testing.T) {
	file, err := os.Open("data.txt")
	assert.NoError(t, err)
	defer file.Close()

	beefCounts := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value, err := strconv.Atoi(strings.TrimSpace(parts[1]))
			assert.NoError(t, err)
			beefCounts[key] = value
		}
	}
	assert.NoError(t, scanner.Err())

	// Create a new Fiber app and register the handler
	app := fiber.New()
	app.Get("/beef/summary", GetBeefSummary)

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/beef/summary", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	// Check HTTP status
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Parse response JSON
	var response map[string]int
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)

	// Check if the response matches the expected beefCounts
	for key, value := range beefCounts {
		assert.Equal(t, value, response[key])
	}
	for key, value := range response {
		assert.Equal(t, value, beefCounts[key])
	}
}
