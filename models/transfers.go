package models

// Stores the data received by the query.
type Transfers struct {
	Count    int64
	Days     int
	Min_date string
	Max_date string
}
