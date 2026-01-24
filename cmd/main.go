package main

import (
	"fmt"
	"bible/services"
)

func main() {
	bibleService := services.NewBibleService()
	dailyReading, err := bibleService.GetDailyReading(1)
	if err != nil {
		fmt.Printf("Error getting daily reading: %v\n", err)
		return
	}
	services.PrintDailyReading(dailyReading)
}
