package conversion

import (
	apiModels "hu.erettsegizteto/internal/api/models"
	dbModels "hu.erettsegizteto/internal/db/models"
)

func DBQuestionToAPIQuestion(question *dbModels.Question) *apiModels.Question {
	if question == nil {
		return nil
	}

	return &apiModels.Question{
		ID:            question.ID,
		Question:      question.Question,
		Generated:     question.Generated,
		Year:          question.Year,
		Semester:      question.Semester,
		Number:        question.Number,
		AnswerHolders: DBAnswerHoldersToAPIAnswerHolders(question.AnswerHolders),
	}
}

func DBAnswerHolderToAPIAnswerHolder(answerHolder dbModels.AnswerHolder) apiModels.AnswerHolder {
	return apiModels.AnswerHolder{
		ID:         answerHolder.ID,
		QuestionID: answerHolder.QuestionID,
		Prefix:     answerHolder.Prefix,
		Suffix:     answerHolder.Suffix,
		Number:     answerHolder.Number,
		Help:       answerHolder.Help,
	}
}

func DBAnswerHoldersToAPIAnswerHolders(answerHolders []dbModels.AnswerHolder) []apiModels.AnswerHolder {
	apiAhs := []apiModels.AnswerHolder{}

	for _, dbAh := range answerHolders {
		apiAhs = append(apiAhs, DBAnswerHolderToAPIAnswerHolder(dbAh))
	}

	return apiAhs
}
