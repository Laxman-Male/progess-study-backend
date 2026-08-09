package mcqtest

import (
	"database/sql"
	"errors"
	"strings"

	allstruct "progress_tracker/allStruct"
	"progress_tracker/database"
	"progress_tracker/queries"
)

// SubmitQuestionRepo grades the answer for mcqID, stores the selected option,
// marks it attempted, and returns the next un-attempted MCQ plus the overall
// attempted/total count. A zero-value MCQFormat (MCQ == "") with a nil error
// means the quiz is complete.
func SubmitQuestionRepo(optionSelected, email, title string, mcqID int) (allstruct.MCQFormat, allstruct.MCQCounts, error) {

	var UserId string
	var planID int
	var oneMcq allstruct.MCQFormat
	var crrAns string

	err := database.DB.Get(&UserId, queries.GetUserIDForProfile, email)
	if err != nil {
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	err = database.DB.Get(&planID, queries.GetPlanIDForParticularPlan, UserId, title)
	if err != nil {
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	err = database.DB.Get(&crrAns, queries.ValidatingAns, planID, mcqID)
	if err != nil {
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	_, err = database.DB.Exec(queries.UpdateOptionSelected, optionSelected, planID, mcqID)
	if err != nil {
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	// crrAns may be lowercase for plans saved before answers were normalized to uppercase.
	if strings.EqualFold(optionSelected, crrAns) {
		_, err = database.DB.Exec(queries.UpdateAnsStateAsTrue, planID, mcqID)
	} else {
		_, err = database.DB.Exec(queries.UpdateAnsStateAsFalse, planID, mcqID)
	}
	if err != nil {
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	_, err = database.DB.Exec(queries.UpdateAttemptedState, planID, mcqID)
	if err != nil {
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	counts, err := getMCQCounts(planID)
	if err != nil {
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	err = database.DB.Get(&oneMcq, queries.GetOneMcqTosolve, planID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return allstruct.MCQFormat{}, counts, nil
		}
		return allstruct.MCQFormat{}, allstruct.MCQCounts{}, err
	}

	return oneMcq, counts, nil
}
