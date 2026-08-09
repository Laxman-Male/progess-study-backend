package repo

import (
	"encoding/json"
	"fmt"
	allstruct "progress_tracker/allStruct"
	"progress_tracker/database"
	"progress_tracker/queries"
	"strings"
	"time"
)

func Saveplan(userPlan allstruct.SaveUserPlanToDB, userEmail string) error {
	fmt.Println("in repo")
	var userId string
	var WeekCount int = userPlan.WeekCount
	var title string = userPlan.Title
	cleanedTitle := strings.Trim(title, "\"")
	fmt.Println("cleaned", cleanedTitle)

	strPlan, err := json.Marshal(userPlan.Plan)        //marshal send []byte array means [as below]
	fmt.Println("------------------userplan", strPlan) //this print ACCII value of JSON string
	fmt.Println("----------userplan", string(strPlan))
	if err != nil {
		return err
	}
	err = database.DB.Get(&userId, queries.GetIdByEmailToSavePlan, userEmail)
	if err != nil {
		fmt.Println("error in save repo", err)
	}
	fmt.Println("save repo id= ", userId)

	createdAt := time.Now().UnixMilli()
	res, err := database.DB.Exec(queries.SavePlanOfUser, strPlan, userId, cleanedTitle, WeekCount, createdAt)
	if err != nil {
		fmt.Println("error in exec to save plan")
		fmt.Println(err)
		return err
	}

	planID, err := res.LastInsertId()
	if err != nil {
		fmt.Println("error getting new planID")
		return err
	}

	for _, mcq := range userPlan.MCQs {
		optionsJSON, err := json.Marshal(mcq.Options)
		if err != nil {
			return err
		}
		// Gemini sometimes returns the correct answer letter in lowercase; the
		// options are always keyed A-D, so normalize to keep them matching.
		correctAns := strings.ToUpper(strings.TrimSpace(mcq.CorrectAns))
		_, err = database.DB.Exec(queries.InsertMCQ, mcq.Question, optionsJSON, correctAns, planID)
		if err != nil {
			fmt.Println("error inserting mcq", err)
			return err
		}
	}

	return nil
}
