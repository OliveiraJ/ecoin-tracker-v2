package models

type Read struct {
	BurnedTokens float64 `json:"burnedTokens"`
	Holders      int64   `json:"holders"`
	Transfers    int64   `json:"transfers"`
	Date         string  `json:"date"`
	Hour         string  `json:"hour"`
}
