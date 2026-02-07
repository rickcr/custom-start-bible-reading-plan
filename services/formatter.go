package services

import (
	"bible/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// FormatReading formats a single reading (book + chapters) as a string.
// Example: "Genesis 1-3" or "Matthew 2, 6, 12-14"
func FormatReading(r *model.Reading) string {
	if len(r.Chapters) == 0 {
		return r.Book
	}
	var sb strings.Builder
	sb.WriteString(r.Book)
	for i, ch := range r.Chapters {
		if i == 0 {
			sb.WriteString(" ")
		} else {
			sb.WriteString(", ")
		}
		if ch.End != 0 {
			fmt.Fprintf(&sb, "%d-%d", ch.Start, ch.End)
		} else {
			fmt.Fprintf(&sb, "%d", ch.Start)
		}
	}
	return sb.String()
}

// FormatDay formats all readings for a day as a semicolon-separated string.
// Example: "Genesis 1-3; Titus; Matthew 2, 6, 12-14"
func FormatDay(dr *model.DayReadings) string {
	var sb strings.Builder
	for i, reading := range dr.Readings {
		if i > 0 {
			sb.WriteString("; ")
		}
		sb.WriteString(FormatReading(&reading))
	}
	return sb.String()
}

// WriteReading writes a single reading to the given writer.
func WriteReading(w io.Writer, r *model.Reading) {
	fmt.Fprint(w, FormatReading(r))
}

// WriteDay writes all readings for a day to the given writer.
func WriteDay(w io.Writer, dr *model.DayReadings) {
	fmt.Fprint(w, FormatDay(dr))
}

// WriteAllDays writes all 365 daily readings to the given writer.
func WriteAllDays(w io.Writer, days []model.DayReadings) {
	for _, dr := range days {
		fmt.Fprintf(w, "Day %d: ", dr.Day)
		WriteDay(w, &dr)
		fmt.Fprintln(w)
	}
}

// LoadReadings loads daily readings from a JSON file.
func LoadReadings(filePath string) ([]model.DayReadings, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var readings []model.DayReadings
	if err := json.Unmarshal(data, &readings); err != nil {
		return nil, err
	}

	return readings, nil
}

func GetReadingsForDay(day int, days []model.DayReadings) model.DayReadings {
	idx := day - 1
	if idx < 0 || idx >= len(days) {
		return model.DayReadings{}
	}
	return days[idx]
}
