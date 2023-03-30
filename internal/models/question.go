package models

// Question represents a question in the database
type Question struct {
	ID        int    `json:"id"`
	Question  string `json:"question"`
	Answer    string `json:"answer"`
	Generated bool   `json:"generated"`
	Year      int    `json:"year"`
	Semester  int    `json:"semester"`
	Number    int    `json:"number"`
}
