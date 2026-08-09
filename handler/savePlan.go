package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	// allstruct "progress_tracker/allStruct"
	allstruct "progress_tracker/allStruct"
	"progress_tracker/service"
)

//need to change 	service.SavePlanSer(USplan, userEmail) this line to USplan.plan i just want to store the plan without any count in planSavedByUser

func SavePlanOfUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var USplan allstruct.SaveUserPlanToDB
		// var USplan any
		err := json.NewDecoder(r.Body).Decode(&USplan)
		var weekCountOfPlan = USplan.WeekCount
		var title = USplan.Title
		fmt.Println("title & w count", title, weekCountOfPlan)
		// var rawJSON json.RawMessage
		// err := json.Unmarshal([]byte(USplan), &rawJSON)
		userEmail, ok := GetUserIDFromContext(r.Context())
		if !ok {
			fmt.Println("in savepaln.go", err)
		}

		fmt.Println("in saves handler")
		if err != nil {
			fmt.Println("error in saving plan")
			http.Error(w, "Error decoding plan data", http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		fmt.Println("US plan handler", USplan)
		err = service.SavePlanSer(USplan, userEmail)
		if err != nil {
			fmt.Println("error saving plan", err)
			http.Error(w, "error saving plan", http.StatusInternalServerError)
			return
		}

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

}
