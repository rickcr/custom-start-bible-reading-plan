package services

import (
	"bible/model"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type BibleService struct {
	// Add any necessary fields here, such as a database connection or API client
}

func NewBibleService() *BibleService {
	return &BibleService{}
}

func (bs *BibleService) GetDailyReading(day int) (*model.DayReadings, error) {
	dailyReading := &model.DayReadings{
		Day: day,
		Readings: []model.Reading{
			{Book: "Genesis", Chapters: []model.Chapter{
				{Start: 1, End: 3},
			}},
			{Book: "Titus"},
			{Book: "Mathew", Chapters: []model.Chapter{
				{Start: 2}, {Start: 6}, {Start: 12, End: 14},
			}},
		},
	}
	return dailyReading, nil
}

func DailyReadingsOutput(dailyReading *model.DayReadings) string {
	var buf bytes.Buffer
	for i, reading := range dailyReading.Readings {
		if i > 0 {
			buf.WriteString("; ")
		}
		buf.WriteString(DailyReadingOutput(&reading))
	}
	return buf.String()
}

func PrintDailyReadings(dailyReading *model.DayReadings) {
	for i, reading := range dailyReading.Readings {
		if i > 0 {
			fmt.Print("; ")
		}
		PrintReadingEntry(&reading)
	}
}

func PrintAllDailyReadings(dailyReadings []model.DayReadings) {
	for _, dr := range dailyReadings {
		fmt.Printf("Day %d: ", dr.Day)
		PrintDailyReadings(&dr)
		fmt.Println()
	}
}

func DailyReadingOutput(reading *model.Reading) string {
	var buf bytes.Buffer
	buf.WriteString(reading.Book)
	for i, chapter := range reading.Chapters {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(" ")
		buf.WriteString(fmt.Sprintf("%d", chapter.Start))
		if chapter.End != 0 {
			buf.WriteString(fmt.Sprintf("-%d", chapter.End))
		}
	}
	return buf.String()
}

func PrintReadingEntry(reading *model.Reading) {
	fmt.Print(DailyReadingOutput(reading))
}

func LoadChronoReadings(filePath string) ([]model.DayReadings, error) {
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
