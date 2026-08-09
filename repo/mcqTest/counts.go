package mcqtest

import (
	allstruct "progress_tracker/allStruct"
	"progress_tracker/database"
	"progress_tracker/queries"
)

// getMCQCounts returns how many MCQs exist for the plan and how many have been attempted.
func getMCQCounts(planID int) (allstruct.MCQCounts, error) {
	var counts allstruct.MCQCounts
	err := database.DB.Get(&counts, queries.GetMcqCounts, planID)
	if err != nil {
		return allstruct.MCQCounts{}, err
	}
	return counts, nil
}
