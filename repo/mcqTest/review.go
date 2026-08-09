package mcqtest

import (
	allstruct "progress_tracker/allStruct"
	"progress_tracker/database"
	"progress_tracker/queries"
)

// GetPlanReviewRepo returns all MCQs for the user's plan, including the
// selected option and whether it was correct, for the post-quiz review screen.
func GetPlanReviewRepo(title, email string) ([]allstruct.MCQReview, error) {

	var UserId string
	var planID int
	var review []allstruct.MCQReview

	err := database.DB.Get(&UserId, queries.GetUserIDForProfile, email)
	if err != nil {
		return nil, err
	}

	err = database.DB.Get(&planID, queries.GetPlanIDForParticularPlan, UserId, title)
	if err != nil {
		return nil, err
	}

	err = database.DB.Select(&review, queries.GetAllMcqsForPlan, planID)
	if err != nil {
		return nil, err
	}

	return review, nil
}
