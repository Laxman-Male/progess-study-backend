package service

import (
	"fmt"
	allstruct "progress_tracker/allStruct"
	"progress_tracker/repo"
)

func SavePlanSer(USplan allstruct.SaveUserPlanToDB, userEmail string) error {
	fmt.Println("in service")
	err := repo.Saveplan(USplan, userEmail)
	if err != nil {
		fmt.Println("error in repo to save file")
		return err
	}

	return nil
}
