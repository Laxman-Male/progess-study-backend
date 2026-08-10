package service

import (
	"errors"
	"fmt"
	allstruct "progress_tracker/allStruct"
	"progress_tracker/repo"
)

func ValidateReg(reg allstruct.RegInfo) (string, error) {
	fmt.Println("enter into  service")
	if reg.Name == "" || reg.Password == "" {
		return " ", errors.New("name & password required")
	}
	email, err := repo.ExecuteQniqueEmail(reg.Email)
	if err != nil {
		fmt.Println("you didn't got unique email")
		return "", err
	}
	fmt.Println("service email unique ", email)
	// reg.Email=email;
	if len(reg.Password) < 5 {
		return " ", errors.New("password length must be greater than 5")
	}
	_, err = repo.ExecuteRegQuery(reg)
	if err != nil {
		fmt.Println("error in insert vlaue service")
		return " ", errors.New("error in insert value")
	}

	return reg.Email, nil
}

func ValidateLogin(login allstruct.LoginFromUser) (string, error) {

	fmt.Println("enter login service")
	fmt.Println("before login", login.Email, login.Name)
	var dbDetails allstruct.LoginInfo
	dbDetails, err := repo.ExecuteLoginQuery(login)
	if err != nil {
		return dbDetails.Id, errors.New("error in login in service file")

	}
	fmt.Println("after return", login.Email, login.Name)
	if dbDetails.Email == login.Email && dbDetails.Name == login.Name {
		// if dbDetails.Name == login.Name {
		fmt.Println("login id", dbDetails.Id)
		err := repo.UpdateLogin(dbDetails.Id)
		if err != nil {
			return dbDetails.Id, errors.New("error in login ")
		}
		// fmt.Println("sucees", a)
		fmt.Println("login success ✅")
		//here i can get all the data once user login and can display on web
		// getting it from tables
		return dbDetails.Email, nil
		// } else {
		// 	fmt.Println("password ❌")
		// 	return "", err
		// }
	} else {
		fmt.Println("name ans pass ❌")
		fmt.Println("password ❌")
		// return "", err
		return dbDetails.Id, errors.New("both are incorrect")
	}
	// return dbDetails.Id, nil

}

func LogOut(email, password string) error {
	err := repo.ExecuteLogOut(email, password)
	if err != nil {
		fmt.Println("error in log out service")
		return err
	}
	return nil
}

// func UpdateLogin() error {
// 	err := repo.UpdateLogin(ISLOGIN)
// 	if err != nil{
// 		errors.New("check name ans password")
// 	}
// 	return nil
// }
