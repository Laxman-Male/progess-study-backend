package repo

import (
	"errors"
	"fmt"

	// allstruct "progress_tracker/allStruct"
	"progress_tracker/database"
	"progress_tracker/queries"
	// "github.com/golang/protobuf/jsonpb"
	// "github.com/golang/protobuf/jsonpb"
)

// err := database.DB.Get(&loginInfo, queries.LoginQuery, login.Name)

type values struct {
	Id       int    `db:"id" json:"id"`
	Jsondata string `db:"study" json:"study"`
}

func ExecutePlanQuery(c string) error {
	fmt.Println("enter login repo")
	rr, err := database.DB.Exec(queries.InsertStudyPlan, 1, c)
	if err != nil {
		return errors.New("failed to execute login")
	}

	fmt.Println("after exc of plan", rr)
	// var rawJSON json.RawMessage
	// err = json.Unmarshal([]byte(rr), &rawJSON)

	return nil

}

func GetPlanrepo(id int) (values, error) {

	var a values
	err := database.DB.Get(&a, queries.GetPlan, 1)
	if err != nil {
		fmt.Println("error in getting from db")
	}
	return a, nil

}
