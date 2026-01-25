package services

import (
	"bible/model"
	"bytes"
	"io"
	"os"
	"testing"
)

// captureOutput captures stdout during function execution
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestPrintReadingEntry(t *testing.T) {
	tests := []struct {
		name     string
		reading  model.Reading
		expected string
	}{
		{
			name:     "single chapter",
			reading:  model.Reading{Book: "Genesis", Chapters: []model.Chapter{{Start: 1}}},
			expected: "Genesis 1",
		},
		{
			name:     "chapter range",
			reading:  model.Reading{Book: "Genesis", Chapters: []model.Chapter{{Start: 1, End: 3}}},
			expected: "Genesis 1-3",
		},
		{
			name:     "multiple chapters",
			reading:  model.Reading{Book: "Mathew", Chapters: []model.Chapter{{Start: 2}, {Start: 6}, {Start: 12, End: 14}}},
			expected: "Mathew 2, 6, 12-14",
		},
		{
			name:     "book only",
			reading:  model.Reading{Book: "Titus"},
			expected: "Titus",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				PrintReadingEntry(&tt.reading)
			})
			t.Logf("Output: %q", output)
			if output != tt.expected {
				t.Errorf("got %q, want %q", output, tt.expected)
			}
		})
	}
}

func TestPrintDailyReading(t *testing.T) {
	dailyReadings := &model.DayReadings{
		Day: 1,
		Readings: []model.Reading{
			{Book: "Genesis", Chapters: []model.Chapter{{Start: 1, End: 3}}},
			{Book: "Titus"},
			{Book: "Mathew", Chapters: []model.Chapter{{Start: 2}, {Start: 6}, {Start: 12, End: 14}}},
		},
	}

	output := captureOutput(func() {
		PrintDailyReadings(dailyReadings)
	})

	t.Logf("Output: %q", output)

	expected := "Genesis 1-3; Titus; Mathew 2, 6, 12-14"
	if output != expected {
		t.Errorf("got %q, want %q", output, expected)
	}
}

func TestLoadChronoReadings(t *testing.T) {
	dailyReadings, err := LoadChronoReadings("../data/chrono.json")

	if err != nil {
		t.Fatalf("Failed to load chrono readings: %v", err)
	}

	if len(dailyReadings) == 0 {
		t.Fatalf("Expected non-empty daily readings")
	}

	// Simple check for the first day's readings
	firstDay := dailyReadings[0]
	if firstDay.Day != 1 {
		t.Errorf("Expected day 1, got %d", firstDay.Day)
	}

	if len(firstDay.Readings) == 0 {
		t.Errorf("Expected readings for day 1, got none")
	}

	PrintAllDailyReadings(dailyReadings)
}
