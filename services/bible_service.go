package services

import (
	"bible/model"
	"fmt"
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

func PrintDailyReading(dailyReading *model.DayReadings) {
	for i, reading := range dailyReading.Readings {
		if i > 0 {
			fmt.Print("; ")
		}
		PrintReadingEntry(&reading)
	}
}

func PrintReadingEntry(reading *model.Reading) {
	fmt.Print(reading.Book)
	for i, chapter := range reading.Chapters {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(" ", chapter.Start)
		if chapter.End != 0 {
			fmt.Print("-", chapter.End)
		}
	}
}
