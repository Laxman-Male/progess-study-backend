package service

import (
	"fmt"
	"progress_tracker/repo"
)

func OwnPlanDescriptionService(lowerTitle string, userEmail string) (string, error) {
	plan, err := repo.OwnPlanDescriptionRepo(lowerTitle, userEmail)
	if err != nil {
		return "", err
	}
	return plan, nil
}

func GetWeekPlanViewPlan(email string, title string) (int, error) {

	count, err := repo.GetWeekPlanViewPlan(email, title)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func WeekCompletedService(title, userEmail string, nowTime int64) (int, int, error) {
	completedWeeks, totalWeeks, err := repo.WeekCompletedRepo(title, userEmail, nowTime)
	if err != nil {
		fmt.Println("error in ser/*/*/*")
		return 0, 0, err
	}
	return completedWeeks, totalWeeks, nil
}
