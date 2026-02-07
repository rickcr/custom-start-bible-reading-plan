package main

import (
	"bible/services"
	"fmt"
	"os"
)

func main() {
	days, err := services.LoadReadings("data/chrono.json")
	if err != nil {
		fmt.Printf("Error loading readings: %v\n", err)
		return
	}

	day := services.GetReadingsForDay(1, days)
	services.WriteDay(os.Stdout, &day)
	fmt.Println()
}
