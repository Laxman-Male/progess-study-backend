package mcqtest

import (
	"database/sql"
	"errors"

	allstruct "progress_tracker/allStruct"
	"progress_tracker/database"
	"progress_tracker/queries"
)

// FirstQuestionRepo returns the first un-attempted MCQ for the user's plan, plus
// the overall attempted/total count. A zero-value MCQFormat (MCQ == "") with a
// nil error means all MCQs are done.
func FirstQuestionRepo(title, email string) (allstruct.MCQFormat, allstruct.MCQCounts, error) {

	var UserId string
	var planID int
	var Fmcq allstruct.MCQFormat

	err := database.DB.Get(&UserId, queries.GetUserIDForProfile, email)
	if err != nil {
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	err = database.DB.Get(&planID, queries.GetPlanIDForParticularPlan, UserId, title)
	if err != nil {
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	counts, err := getMCQCounts(planID)
	if err != nil {
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	err = database.DB.Get(&Fmcq, queries.GetOneMcqTosolve, planID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return allstruct.MCQFormat{}, counts, nil
		}
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	return Fmcq, counts, nil
}
