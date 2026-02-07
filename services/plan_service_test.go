package services

import (
	"os"
	"testing"
)

func TestFindFirstDay(t *testing.T) {
	data, err := os.ReadFile("../data/year_chrono.json")
	if err != nil {
		t.Fatalf("failed to read year_chrono.json: %v", err)
	}

	tests := []struct {
		name    string
		book    string
		wantDay int
		wantErr bool
	}{
		{name: "Genesis is day 1", book: "Genesis", wantDay: 1},
		{name: "Job is day 4", book: "Job", wantDay: 4},
		{name: "Exodus is day 30", book: "Exodus", wantDay: 30},
		{name: "Psalms is day 81", book: "Psalms", wantDay: 81},
		{name: "case insensitive", book: "genesis", wantDay: 1},
		{name: "not found", book: "FakeBook", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindFirstDay(data, tt.book)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.Day != tt.wantDay {
				t.Errorf("got day %d, want %d", got.Day, tt.wantDay)
			}
		})
	}
}

func TestFindFirstDay_InvalidJSON(t *testing.T) {
	_, err := FindFirstDay([]byte("not json"), "Genesis")
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}
