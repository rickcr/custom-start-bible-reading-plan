package model

type Chapter struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type Reading struct {
	Book     string  `json:"book"`
	Chapters []Chapter `json:"chapter,omitempty"`
}

type DayReadings struct {
	Day         int     `json:"day"`
	Readings    []Reading `json:"readings"`
}

