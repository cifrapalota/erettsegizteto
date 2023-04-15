package models

type Solution struct {
	CheckedAnswers []CheckedAnswerHolder `json:"checkedAnswers"`
	Solution       string                `json:"solution"`
}
