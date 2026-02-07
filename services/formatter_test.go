package services

import (
	"bible/model"
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestFormatReading(t *testing.T) {
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
			reading:  model.Reading{Book: "Matthew", Chapters: []model.Chapter{{Start: 2}, {Start: 6}, {Start: 12, End: 14}}},
			expected: "Matthew 2, 6, 12-14",
		},
		{
			name:     "book only",
			reading:  model.Reading{Book: "Titus"},
			expected: "Titus",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatReading(&tt.reading)
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestFormatDay(t *testing.T) {
	day := &model.DayReadings{
		Day: 1,
		Readings: []model.Reading{
			{Book: "Genesis", Chapters: []model.Chapter{{Start: 1, End: 3}}},
			{Book: "Titus"},
			{Book: "Matthew", Chapters: []model.Chapter{{Start: 2}, {Start: 6}, {Start: 12, End: 14}}},
		},
	}

	got := FormatDay(day)
	expected := "Genesis 1-3; Titus; Matthew 2, 6, 12-14"
	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}

func TestWriteDay(t *testing.T) {
	day := &model.DayReadings{
		Day: 1,
		Readings: []model.Reading{
			{Book: "Genesis", Chapters: []model.Chapter{{Start: 1, End: 3}}},
			{Book: "Titus"},
		},
	}

	var buf bytes.Buffer
	WriteDay(&buf, day)

	got := buf.String()
	expected := "Genesis 1-3; Titus"
	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}

func TestWriteAllDays(t *testing.T) {
	days := []model.DayReadings{
		{Day: 1, Readings: []model.Reading{{Book: "Genesis", Chapters: []model.Chapter{{Start: 1}}}}},
		{Day: 2, Readings: []model.Reading{{Book: "Genesis", Chapters: []model.Chapter{{Start: 2}}}}},
	}

	var buf bytes.Buffer
	WriteAllDays(&buf, days)

	expected := "Day 1: Genesis 1\nDay 2: Genesis 2\n"
	got := buf.String()
	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}

func TestLoadReadings(t *testing.T) {
	days, err := LoadReadings("../data/chrono.json")
	if err != nil {
		t.Fatalf("Failed to load readings: %v", err)
	}

	if len(days) == 0 {
		t.Fatal("Expected non-empty daily readings")
	}

	firstDay := days[0]
	if firstDay.Day != 1 {
		t.Errorf("Expected day 1, got %d", firstDay.Day)
	}

	if len(firstDay.Readings) == 0 {
		t.Error("Expected readings for day 1, got none")
	}

	WriteAllDays(os.Stdout, days)
}

func TestGetReadingsForDay(t *testing.T) {
	type args struct {
		day                  int
		readingsForDaysArray []model.DayReadings
	}
	tests := []struct {
		name string
		args args
		want model.DayReadings
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReadingsForDay(tt.args.day, tt.args.readingsForDaysArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReadingsForDay() = %v, want %v", got, tt.want)
			}
		})
	}
}
