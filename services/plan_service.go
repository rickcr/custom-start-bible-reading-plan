package services

import (
	"bible/model"
	"encoding/json"
	"fmt"
	"strings"
)

// FindFirstDay takes raw JSON data and a book name, parses the JSON into
// []DayReadings, and returns the first day where that book appears as a reading.
func FindFirstDay(data []byte, book string) (model.DayReadings, error) {
	var days []model.DayReadings
	if err := json.Unmarshal(data, &days); err != nil {
		return model.DayReadings{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	target := strings.ToLower(book)
	for _, day := range days {
		for _, r := range day.Readings {
			if strings.ToLower(r.Book) == target {
				return day, nil
			}
		}
	}

	return model.DayReadings{}, fmt.Errorf("book %q not found in readings", book)
}
