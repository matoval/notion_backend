package main

import (
	"notion_backend/db"
	"notion_backend/routes"
	"notion_backend/utils"
)

func main() {
	utils.LoadEnvVars()
	routes.GetNotionData()
	db.CreateDatabase()
}
