package models

// Works as a adapter so the received data, once this has a particular structure that demands the data to be wraped by other structs
// dificulting it conversion.
type BitqueryData struct {
	Data struct {
		Ethereum struct {
			Transfers []Transfers
		}
	}
}
