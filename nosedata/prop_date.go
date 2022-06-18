package nosedata

import "time"

type DateProp struct {
	Object string `json:"object,omitempty"`
	ID     string `json:"id,omitempty"`
	Type   string `json:"type,omitempty"`
	Date   Date   `json:"date"`
}

type Date struct {
	Start    time.Time  `json:"start"`
	End      *time.Time `json:"end,omitempty"`
	TimeZone *string    `json:"time_zone,omitempty"`
}

func (n *DateProp) PropType() string {
	return "date"
}

func NewDBDateColumn(d time.Time) *DateProp {
	return &DateProp{Date: Date{Start: d}}
}

//for database prop init
type EmptyDatabaseDateProp struct {
	Date struct{} `json:"date"`
}

func (n EmptyDatabaseDateProp) PropType() string {
	return "date"
}
