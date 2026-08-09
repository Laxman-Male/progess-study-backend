package service

import (
	"fmt"
	"progress_tracker/repo"
)

func GetUserPlan(email string) (int, []string, []bool, error) {
	fmt.Println("enter servie to get count")
	count, title, completed, err := repo.GetPlanForProfile(email)
	if err != nil {
		fmt.Println("err ser", err)
		return 0, nil, nil, err
	}
	fmt.Println("count--- ser", count)
	fmt.Println("title----ser", title)
	return count, title, completed, nil
}
