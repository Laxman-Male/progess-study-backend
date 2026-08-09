package repo

import (
	"fmt"
	allstruct "progress_tracker/allStruct"
	"progress_tracker/database"
	"progress_tracker/queries"
)

func Profile(userEmail string) (allstruct.GetProfileDetails, error) {
	var user allstruct.GetProfileDetails
	fmt.Println("userid=", userEmail)
	// fmt.Println("userid=",type)
	err := database.DB.Get(&user, queries.GetUserDetail, userEmail)
	if err != nil {
		fmt.Printf("error in getting from db profile=%v", err)
		return user, err
	}
	return user, nil

}
