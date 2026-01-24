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
	dailyReading := &model.DayReadings{
		Day: 1,
		Readings: []model.Reading{
			{Book: "Genesis", Chapters: []model.Chapter{{Start: 1, End: 3}}},
			{Book: "Titus"},
			{Book: "Mathew", Chapters: []model.Chapter{{Start: 2}, {Start: 6}, {Start: 12, End: 14}}},
		},
	}

	output := captureOutput(func() {
		PrintDailyReading(dailyReading)
	})

	t.Logf("Output: %q", output)

	expected := "Genesis 1-3; Titus; Mathew 2, 6, 12-14"
	if output != expected {
		t.Errorf("got %q, want %q", output, expected)
	}
}
