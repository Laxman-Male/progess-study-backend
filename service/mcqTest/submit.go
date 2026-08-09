package mcqtest

import (
	"fmt"
	allstruct "progress_tracker/allStruct"
	mcqtest "progress_tracker/repo/mcqTest"
)

func SubmitQuestionService(optionSelected string, email, title string, mcqID int) (allstruct.MCQFormat, allstruct.MCQCounts, error) {

	mcq, counts, err := mcqtest.SubmitQuestionRepo(optionSelected, email, title, mcqID)
	if err != nil {
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}
	fmt.Println("mcq in service", mcq)

	return mcq, counts, nil
}
