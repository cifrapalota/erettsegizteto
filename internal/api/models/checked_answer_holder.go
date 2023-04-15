package models

type CheckedAnswerHolder struct {
	AnswerHolderID string   `json:"answer_holder_id"`
	Correct        bool     `json:"correct"`
	Answers        []string `json:"answers"`
}
