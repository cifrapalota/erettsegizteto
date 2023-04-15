package models

type CheckedAnswerHolder struct {
	AnswerHolderID string   `json:"answerHolderId"`
	Correct        bool     `json:"correct"`
	Answers        []string `json:"answers"`
}
