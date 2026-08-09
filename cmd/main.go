package main

import (
	"log"
	"net/http"
	"progress_tracker/config"
	"progress_tracker/database"
	"progress_tracker/routes"
)

func main() {
	cfg := config.LoadConfig()
	database.Connect(cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBName)
	// apiKey := os.Getenv("geminiApi")
	// if apiKey==""{
	// 	fmt.Println("failed to connect with API")
	// 	return
	// }

	// Run DB migration when table is not present
	// models.Migrate(database.DB)

	routes := routes.SetUpRoutes()
	log.Println("server started at 8000")
	err := http.ListenAndServe(":8000", routes)
	if err != nil {
		log.Fatal("failed to start", err)
	}

}
