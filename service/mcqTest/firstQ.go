package mcqtest

import (
	"fmt"
	allstruct "progress_tracker/allStruct"
	mcqtest "progress_tracker/repo/mcqTest"
)

func FirstQuestionService(title, email string) (allstruct.MCQFormat, allstruct.MCQCounts, error) {

	Fmcq, counts, err := mcqtest.FirstQuestionRepo(title, email)
	if err != nil {
		fmt.Println("err in getting first Q")
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}
	return Fmcq, counts, nil
}
