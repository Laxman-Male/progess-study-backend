package repo

import (
	"database/sql"
	"errors"
	"fmt"

	"progress_tracker/database"
	"progress_tracker/queries"
)

type plan struct {
	// Count int    `db:"count"`
	// Title string `db:"title"`
	Plan string `db:"planSavedByUser"`
}

func OwnPlanDescriptionRepo(title string, userEmail string) (string, error) {
	fmt.Println("in repo to get description")
	var ID string
	// var plan []string
	var plan plan
	// var titleList string
	err := database.DB.Get(&ID, queries.GetIdByEmailToSavePlan, userEmail)
	if err != nil {
		return "", fmt.Errorf("error in getting ID")
		// return allstruct.GoStudyPlanWeekGapOne{}, fmt.Errorf("error in getting ID")
	}

	err = database.DB.Get(&plan, queries.GetOwnPlanDescp, ID, title)
	if err != nil {
		fmt.Println("err in plan desp", err)
	}
	fmt.Println("title in desp", plan)
	// var parsedPlan allstruct.GoStudyPlanWeekGapOne
	// err = json.Unmarshal([]byte(plan.Plan), &parsedPlan)
	// if err != nil {
	// 	return allstruct.GoStudyPlanWeekGapOne{}, fmt.Errorf("failed to unmarshal plan JSON: %v", err)
	// }

	// fmt.Println("parsed plan:", parsedPlan)
	// fmt.Println("title in desp", plan.Title)
	// fmt.Println("title in desp", plan.Plan)

	return plan.Plan, nil
}

func GetWeekPlanViewPlan(email string, title string) (int, error) {
	var count int
	var IDByEmail string
	err := database.DB.Get(&IDByEmail, queries.GetIdByEmailToSavePlan, email)
	if err != nil {
		return 0, err
	}
	err = database.DB.Get(&count, queries.GetWeekCountTOshowAtDesp, IDByEmail, title)
	if err != nil {
		return 0, err
	}
	fmt.Println("week count is ----", count)

	return count, nil
}

type planCreatedAndWeekCount struct {
	CreatedAt int64 `db:"CreatedAt"`
	WeekCount int   `db:"weekCount"`
}

// WeekCompletedRepo computes how many weeks of the plan are completed purely from
// CreatedAt vs nowTime (sent by the frontend). No DB write - completedWeek has no
// dedicated tracking column, so this is calculated fresh on every call.
func WeekCompletedRepo(title, userEmail string, nowTime int64) (int, int, error) {

	var userID string
	var info planCreatedAndWeekCount

	err := database.DB.Get(&userID, queries.GetUserIDForProfile, userEmail)
	if err != nil {
		return 0, 0, err
	}

	err = database.DB.Get(&info, queries.GetCreatedAtAndWeekCount, userID, title)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No matching record found")
			return 0, 0, nil
		}
		fmt.Println("error in repo get", err)
		return 0, 0, err
	}

	diff := nowTime - info.CreatedAt
	diffDays := diff / (1000 * 60 * 60 * 24)
	completedWeeks := int(diffDays / 7)

	if completedWeeks < 0 {
		completedWeeks = 0
	}
	if completedWeeks > info.WeekCount {
		completedWeeks = info.WeekCount
	}

	return completedWeeks, info.WeekCount, nil
}
