package repo

import (
	"fmt"
	"time"

	"progress_tracker/database"
	"progress_tracker/queries"
)

type planStatusRow struct {
	Title        string `db:"title"`
	PlanID       int    `db:"planID"`
	CreatedAt    int64  `db:"createdAt"`
	WeekCount    int    `db:"weekCount"`
	TotalMcq     int    `db:"totalMcq"`
	AttemptedMcq int    `db:"attemptedMcq"`
}

// GetPlanForProfile returns the user's plan count, titles, and whether each
// plan is fully completed - all weeks elapsed (based on CreatedAt) AND every
// generated MCQ attempted - in the same order, for the "my plan" list.
func GetPlanForProfile(email string) (int, []string, []bool, error) {
	fmt.Println("getting count repo----")
	var id string
	err := database.DB.Get(&id, queries.GetUserIDForProfile, email)
	if err != nil {
		fmt.Println("err in repo getting count repo----")
		fmt.Println("err", err)
		return 0, nil, nil, err
	}

	var rows []planStatusRow
	err = database.DB.Select(&rows, queries.GetPlanStatusForProfile, id)
	if err != nil {
		fmt.Println("err getting plan status", err)
		return 0, nil, nil, err
	}

	nowTime := time.Now().UnixMilli()
	titleList := make([]string, 0, len(rows))
	completed := make([]bool, 0, len(rows))

	for _, row := range rows {
		titleList = append(titleList, row.Title)

		diffDays := (nowTime - row.CreatedAt) / (1000 * 60 * 60 * 24)
		completedWeeks := min(int(diffDays/7), row.WeekCount)
		weeksDone := row.WeekCount > 0 && completedWeeks >= row.WeekCount
		mcqDone := row.TotalMcq > 0 && row.AttemptedMcq >= row.TotalMcq
		completed = append(completed, weeksDone && mcqDone)
	}

	fmt.Println("count----------repo", len(rows))
	fmt.Println("titleList----------repo", titleList)
	fmt.Println("completed----------repo", completed)

	return len(rows), titleList, completed, nil
}
