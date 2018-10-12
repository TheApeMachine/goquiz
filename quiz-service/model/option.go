package model

// Option : Defines an answer option for a Quiz
type Option struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Correct bool   `json:"correct"`
}
