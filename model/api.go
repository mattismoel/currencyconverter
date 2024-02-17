package model

// Base API response type from frankfurter.app.
type ApiResponse struct {
	Amount float32            `json:"amount,omitempty"` // The amount of base currency.
	Base   string             `json:"base,omitempty"`   // The base currency code.
	Date   string             `json:"date,omitempty"`   // The date. Defaults to "latest".
	Rates  map[string]float32 `json:"rates,omitempty"`  // The map consisting of a given currency code with an associated rate.
}

