package repo

import (
	"errors"
	"fmt"
	allstruct "progress_tracker/allStruct"
	"progress_tracker/database"
	"progress_tracker/queries"

	"github.com/google/uuid"
)

func ExecuteRegQuery(reg allstruct.RegInfo) (string, error) {
	fmt.Println("enter into  repo")

	_, err := database.DB.Exec(queries.RegisterQuery, CreateUserID(), reg.Name, reg.Password, reg.Email)
	if err != nil {
		return " ", errors.New("failed to insert value repo")
	}
	var userId string
	err = database.DB.Get(&userId, queries.GetIdOfUser, reg.Name, reg.Password)
	if err != nil {
		return " ", errors.New("error in getting user id")
	}
	return userId, nil
}

func ExecuteQniqueEmail(inputEmail string) (string, error) {

	fmt.Println("repo for unique email")
	// var emails []string
	var oneEmail string
	//get all email at a time
	rows, err := database.DB.Query(queries.QueryForUniqueEmail)
	if err != nil {
		fmt.Println("getting row error")
	}
	//check if rows are present and close the connection and relese the network resources
	defer rows.Close()
	var allEmail []string
	//rows.Next() to get one email at a time and push into slice
	for rows.Next() {
		err := rows.Scan(&oneEmail)
		if err != nil {
			fmt.Println("scannin error of email")
		}
		allEmail = append(allEmail, oneEmail)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("error in err()")
		return "", err
	}

	fmt.Println("list-", allEmail)
	if len(allEmail) == 0 {
		fmt.Println("No emails fetched from DB")
	}

	for _, value := range allEmail {
		if inputEmail == value {
			// fmt.Println("value in for loop", value)
			return "", errors.New(" not a unique email ")
		}
	}
	return "eeeeemmmail", nil
}

func ExecuteLoginQuery(login allstruct.LoginFromUser) (allstruct.LoginInfo, error) {
	fmt.Println("enter login repo")
	var loginInfo allstruct.LoginInfo
	err := database.DB.Get(&loginInfo, queries.LoginQuery, login.Email)
	if err != nil {
		fmt.Println("failed to execute login query ")
		return loginInfo, errors.New("failed to execute login")
	}

	return loginInfo, nil

}
func UpdateLogin(id string) error {
	_, err := database.DB.Exec(queries.UpdateLoginState, id)
	if err != nil {
		return errors.New("failed to execute login update the login ")

	}

	return nil
}
func ExecuteLogOut(email, password string) error {
	_, err := database.DB.Exec(queries.LogOut, email, password)
	if err != nil {
		fmt.Println("log out in repo")
		return err
	}
	return nil
}
func CreateUserID() string {
	return uuid.New().String()
}

// CREATE TABLE users (
//     id CHAR(36) PRIMARY KEY,
//     name VARCHAR(100),
//     email VARCHAR(100) UNIQUE,
//     password VARCHAR(100)
// );
