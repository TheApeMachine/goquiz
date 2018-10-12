package model

// Step : Defines a step in the Quiz.
type Step struct {
	ID       string `json:"id"`
	Position int    `json:"position"`
	Options  []Option
}
