package main

import (
	"notion_backend/routes"
	"notion_backend/utils"
)

func main() {
	utils.LoadEnvVars()
	routes.GetNotionData()
}
