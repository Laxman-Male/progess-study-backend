package mcqtest

import (
	allstruct "progress_tracker/allStruct"
	mcqtest "progress_tracker/repo/mcqTest"
)

func GetPlanReviewService(title, email string) ([]allstruct.MCQReview, error) {
	return mcqtest.GetPlanReviewRepo(title, email)
}
