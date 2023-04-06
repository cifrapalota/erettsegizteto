package models

// Question represents a question in the database
type Question struct {
	ID        int     `json:"id"`
	Question  string  `json:"question"`
	Answer    string  `json:"answer"`
	Generated bool    `json:"generated"`
	Year      *int    `json:"year,omitempty"`
	Semester  *int    `json:"semester,omitempty"`
	Number    *int    `json:"number,omitempty"`
	ImageLink *string `json:"imageLink,omitempty"`
}
