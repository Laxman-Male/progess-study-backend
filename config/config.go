package config

import (
	"fmt"
	"os"
	allstruct "progress_tracker/allStruct"

	"github.com/joho/godotenv"
)



func LoadConfig() allstruct.DBconnect {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error is featch", err)
	}
	config := allstruct.DBconnect{
		DBUser: os.Getenv("DBUser"),
		DBPass: os.Getenv("DBPass"),
		DBHost: os.Getenv("DBHost"),
		DBName: os.Getenv("DBName"),
	}
	return config

}
