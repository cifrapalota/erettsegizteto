package models

type Workings struct {
	CheckedAnswers []CheckedAnswerHolder `json:"checkedAnswers"`
	Workings       string                `json:"workings"`
}
